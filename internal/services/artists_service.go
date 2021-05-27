package services

import (
	"errors"
	"fmt"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/repository"
	"setmaker-api-go-rest/internal/utils"
	ae "setmaker-api-go-rest/internal/utils/errors"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

type ArtistService struct {
	repository *repository.ArtistsRepository
}

type ValidatorResponse map[string][]string

func NewArtistsService(r *repository.ArtistsRepository) *ArtistService {
	return &ArtistService{
		repository: r,
	}
}

func (svc *ArtistService) GetArtists(filter *utils.QuerySort) ([]*domain.Artist, *ae.AppError) {
	res, err := svc.repository.Find(filter)
	if err != nil {
		return nil, ae.MakeError(ae.ERRInternalServerError, "Error fetching artists")
	}

	if res != nil {
		return res, nil
	}

	return nil, ae.MakeError(ae.ERRNotFound, "No artists found")
}

func (svc *ArtistService) GetArtist(id uuid.UUID) (*domain.Artist, *ae.AppError) {
	artist, err := svc.repository.Get(id)
	if err != nil {
		return nil, ae.MakeError(ae.ERRNotFound, err.Error())
	}

	return artist, nil
}

func (svc *ArtistService) CreateArtist(artist *domain.Artist) *ae.AppError {
	// create unique slug
	err := svc.uniqueSlug(artist)
	if err != nil {
		log.Error("Error generating URL slug")
		return ae.MakeError(ae.ERRInternalServerError, "Could not create artist path")
	}

	// validate
	if errStr := artist.Validate(); len(errStr) > 0 {
		log.Error("Validation error", errStr)
		return ae.MakeError(ae.ERRBadRequest, errStr) // hacky but it'll do until i build a more robust error abstraction
	}

	// create artist
	err = svc.repository.Create(artist)
	if err != nil {
		log.Error("Mongo error:", err)
		return ae.MakeError(ae.ERRInternalServerError, "Could not create artist")
	}

	return nil
}

func (svc *ArtistService) UpdateArtist(a *domain.Artist, id uuid.UUID) *ae.AppError {
	// check the original artist actually exists
	originalArtist, err := svc.GetArtist(id)
	if err != nil {
		return err
	}

	a.ID = id // append ID to artist struct

	// if name is different, the slug will need to changes
	if a.Name != originalArtist.Name {
		err := svc.uniqueSlug(a)
		if err != nil {
			return ae.MakeError(ae.ERRInternalServerError, "Could not create artist path")
		}
	} else {
		// set slug - this doesn't come via the request
		a.Slug = originalArtist.Slug
	}

	// validate the artist struct
	if errStr := a.Validate(); len(errStr) > 0 {
		log.Error("Validation error", errStr)
		return ae.MakeError(ae.ERRBadRequest, errStr)
	}

	e := svc.repository.Update(a)
	if e != nil {
		return ae.MakeError(ae.ERRBadRequest, fmt.Sprintf("Error persisting artist update: %q", id))
	}

	return nil
}

func (svc *ArtistService) DeleteArtist(id uuid.UUID) (*domain.Artist, *ae.AppError) {
	// check the artist exists
	artist, err := svc.GetArtist(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	e := svc.repository.Delete(artist)
	if e != nil {
		return nil, ae.MakeError(ae.ERRInternalServerError, fmt.Sprintf("Artist could not be deleted: %q", id))
	}

	return artist, nil
}

func (svc *ArtistService) uniqueSlug(a *domain.Artist) error {
	// loop through up to n times to create a unique slug
	var s string
	for i := 0; i < 20; i++ {
		if i == 0 {
			s = a.CreateSlug("")
		} else {
			ent := strconv.Itoa(i)
			s = a.CreateSlug(ent)
		}

		// call DB to see if slug is unique
		fs := utils.FieldSearch{
			Field: "slug",
			Query: s,
		}

		count, err := svc.repository.Count(fs)
		if err != nil {
			return err
		}

		// slug is unique, we're done here
		if count == 0 {
			return nil
		}
	}

	return errors.New("Could not generate URL slug for Artist")
}
