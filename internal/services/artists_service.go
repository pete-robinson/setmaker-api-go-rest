package services

import (
	"errors"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/repository"
	"setmaker-api-go-rest/internal/utils"
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

func (svc *ArtistService) GetArtists(filter *utils.QuerySort) ([]*domain.Artist, error) {
	return svc.repository.Find(filter)
}

func (svc *ArtistService) GetArtist(id uuid.UUID) (*domain.Artist, error) {
	return svc.repository.Get(id)
}

func (svc *ArtistService) CreateArtist(artist *domain.Artist) error {
	// create unique slug
	err := svc.uniqueSlug(artist)
	if err != nil {
		log.Error("Error generating URL slug")
		return err
	}

	// validate
	if errStr := artist.Validate(); len(errStr) > 0 {
		log.Error("Validation error", errStr)
		return errors.New(errStr[0]) // hacky but it'll do until i build a more robust error abstraction
	}

	// create artist
	err = svc.repository.Create(artist)
	if err != nil {
		log.Error("Mongo error:", err)
		return err
	}

	return nil
}

func (svc *ArtistService) UpdateArtist(a *domain.Artist, id uuid.UUID) error {
	a.ID = id

	originalArtist, err := svc.repository.Get(id)
	if err != nil {
		log.Error(err)
		return err
	}

	if originalArtist == nil {
		e := errors.New("Artist was not found")
		log.Error(e)
		return e
	}

	if a.Name != originalArtist.Name {
		err := svc.uniqueSlug(a)
		if err != nil {
			return err
		}
	} else {
		a.Slug = originalArtist.Slug
	}

	if errStr := a.Validate(); len(errStr) > 0 {
		log.Error("Validation error", errStr)
		return errors.New(errStr[0])
	}

	return svc.repository.Update(a)
}

func (svc *ArtistService) DeleteArtist(id uuid.UUID) (*domain.Artist, error) {
	// check the artist exists
	artist, err := svc.GetArtist(id)
	if artist == nil || err != nil {
		return nil, errors.New("Artist not found")
	}

	err = svc.repository.Delete(artist)
	if err != nil {
		return nil, err
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
