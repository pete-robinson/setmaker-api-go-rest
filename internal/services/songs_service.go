package services

import (
	"context"
	"fmt"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/repository"
	ae "setmaker-api-go-rest/internal/utils/errors"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type SongService interface {
	GetSongsByArtistId(context.Context, uuid.UUID) ([]*domain.Song, *ae.AppError)
	GetSong(context.Context, uuid.UUID) (*domain.Song, *ae.AppError)
	CreateSong(context.Context, *domain.Song) *ae.AppError
	UpdateSong(context.Context, *domain.Song, uuid.UUID) *ae.AppError
	DeleteSong(context.Context, uuid.UUID) (*domain.Song, *ae.AppError)
}

type songService struct {
	repository repository.SongsRepository
	as         ArtistService
}

// create a new song service
func NewSongsService(r repository.SongsRepository, as ArtistService) *songService {
	return &songService{
		repository: r,
		as:         as,
	}
}

func (svc *songService) GetSongsByArtistId(ctx context.Context, id uuid.UUID) ([]*domain.Song, *ae.AppError) {
	// validate artist exists
	if _, err := svc.as.GetArtist(ctx, id); err != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid artist")
	}

	results := []*domain.Song{}
	results, err := svc.repository.FindSongsByArtistId(ctx, id)
	if err != nil {
		return nil, ae.MakeError(ae.ERRInternalServerError, err)
	}

	return results, nil
}

// Get song
func (svc *songService) GetSong(ctx context.Context, id uuid.UUID) (*domain.Song, *ae.AppError) {
	song, err := svc.repository.GetById(ctx, id)
	if err != nil {
		return nil, ae.MakeError(ae.ERRNotFound, err.Error())
	}

	return song, nil
}

// Create Song
func (svc *songService) CreateSong(ctx context.Context, song *domain.Song) *ae.AppError {
	// validate
	if errStr := song.Validate(); len(errStr) > 0 {
		log.Error("Validation error", errStr)
		return ae.MakeError(ae.ERRBadRequest, errStr)
	}

	// create
	err := svc.repository.Create(ctx, song)
	if err != nil {
		log.Error("Mongo error:", err)
		return ae.MakeError(ae.ERRInternalServerError, "Could not create song")
	}

	return nil
}

func (svc *songService) UpdateSong(ctx context.Context, s *domain.Song, id uuid.UUID) *ae.AppError {
	_, err := svc.GetSong(ctx, id)
	if err != nil {
		return err
	}

	s.ID = id

	if verr := s.Validate(); len(verr) > 0 {
		log.Error("Validation error", verr)
		return ae.MakeError(ae.ERRBadRequest, verr)
	}

	if e := svc.repository.Update(ctx, s); e != nil {
		log.Error(fmt.Sprintf("song update did not save. Err: %v", e))
		return ae.MakeError(ae.ERRBadRequest, fmt.Sprintf("Error persisting song update: %q", id))
	}

	return nil
}

func (svc *songService) DeleteSong(ctx context.Context, id uuid.UUID) (*domain.Song, *ae.AppError) {
	// check the song exists
	song, err := svc.GetSong(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if e := svc.repository.Delete(ctx, song); e != nil {
		return nil, ae.MakeError(ae.ERRInternalServerError, fmt.Sprintf("Song could not be deleted: %q", id))
	}

	return song, nil
}
