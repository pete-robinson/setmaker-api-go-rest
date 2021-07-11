package services

import (
	"context"
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

const MaxSlugIncrement = 10

type ArtistService interface {
	GetArtists(context.Context, *utils.QuerySort) ([]*domain.Artist, *ae.AppError)
	GetArtist(context.Context, uuid.UUID) (*domain.Artist, *ae.AppError)
	CreateArtist(context.Context, *domain.Artist) *ae.AppError
	UpdateArtist(context.Context, *domain.Artist, uuid.UUID) *ae.AppError
	DeleteArtist(context.Context, uuid.UUID) (*domain.Artist, *ae.AppError)
}

type artistService struct {
	repository repository.ArtistsRepository
}

/**
 * instantiate new service
 */
func NewArtistsService(r repository.ArtistsRepository) *artistService {
	return &artistService{
		repository: r,
	}
}

/**
 * Get a list of artists
 */
func (svc *artistService) GetArtists(ctx context.Context, filter *utils.QuerySort) ([]*domain.Artist, *ae.AppError) {
	res, err := svc.repository.Find(ctx, filter)
	if err != nil {
		return nil, ae.MakeError(ae.ERRInternalServerError, "Error fetching artists")
	}

	if res != nil {
		return res, nil
	}

	return nil, ae.MakeError(ae.ERRNotFound, "No artists found")
}

/**
 * Get a single artist by ID
 */
func (svc *artistService) GetArtist(ctx context.Context, id uuid.UUID) (*domain.Artist, *ae.AppError) {
	artist, err := svc.repository.GetById(ctx, id)
	if err != nil {
		return nil, ae.MakeError(ae.ERRNotFound, err.Error())
	}

	return artist, nil
}

/**
 * Create a new artist
 */
func (svc *artistService) CreateArtist(ctx context.Context, artist *domain.Artist) *ae.AppError {
	// create unique slug
	err := svc.uniqueSlug(ctx, artist)
	if err != nil {
		log.Error("Error generating URL slug")
		return ae.MakeError(ae.ERRInternalServerError, "Could not create artist path")
	}

	// validate
	if errStr := artist.Validate(); len(errStr) > 0 {
		log.Error("Validation error", errStr)
		return ae.MakeError(ae.ERRBadRequest, errStr) // hacky but it'll do until i build a more robust error abstraction
	}

	// spawn new UUID
	artist.ID = uuid.New()

	// create artist
	err = svc.repository.Create(ctx, artist)
	if err != nil {
		log.Error("Mongo error:", err)
		return ae.MakeError(ae.ERRInternalServerError, "Could not create artist")
	}

	return nil
}

/**
 * Update an artist
 */
func (svc *artistService) UpdateArtist(ctx context.Context, a *domain.Artist, id uuid.UUID) *ae.AppError {
	// check the original artist actually exists
	originalArtist, err := svc.GetArtist(ctx, id)
	if err != nil {
		return err
	}

	a.ID = id // append ID to artist struct

	// if name is different, the slug will need to changes
	if a.Name != originalArtist.Name {
		err := svc.uniqueSlug(ctx, a)
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

	e := svc.repository.Update(ctx, a)
	if e != nil {
		return ae.MakeError(ae.ERRBadRequest, fmt.Sprintf("Error persisting artist update: %q", id))
	}

	return nil
}

/**
 * Delete an artist
 */
func (svc *artistService) DeleteArtist(ctx context.Context, id uuid.UUID) (*domain.Artist, *ae.AppError) {
	// check the artist exists
	artist, err := svc.GetArtist(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	e := svc.repository.Delete(ctx, artist)
	if e != nil {
		return nil, ae.MakeError(ae.ERRInternalServerError, fmt.Sprintf("Artist could not be deleted: %q", id))
	}

	return artist, nil
}

/**
 * Generate a unique slug from the artist name
 * will query DB to evaluate uniqueness and append an incrementing number if a dupe exists
 * incremented number has a max value of const MaxSlugIncrement
 */
func (svc *artistService) uniqueSlug(ctx context.Context, a *domain.Artist) error {
	// loop through up to n times to create a unique slug
	var s string
	for i := 0; i < MaxSlugIncrement; i++ {
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

		count, err := svc.repository.Count(ctx, fs)
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
