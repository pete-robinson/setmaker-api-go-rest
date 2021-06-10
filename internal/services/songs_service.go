package services

import (
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/repository"
	ae "setmaker-api-go-rest/internal/utils/errors"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type SongService struct {
	repository repository.SongsRepository
}

func (svc *SongService) getRepository() repository.SongsRepository {
	return svc.repository
}

// create a new song service
func NewSongsService(r repository.SongsRepository) *SongService {
	return &SongService{
		repository: r,
	}
}

// Get song
func (svc *SongService) GetSong(id uuid.UUID) (*domain.Song, *ae.AppError) {
	song, err := svc.repository.Get(id)
	if err != nil {
		return nil, ae.MakeError(ae.ERRNotFound, err.Error())
	}

	return song, nil
}

// Create Song
func (svc *SongService) CreateSong(song *domain.Song) *ae.AppError {
	// validate
	if errStr := song.Validate(); len(errStr) > 0 {
		log.Error("Validation error", errStr)
		return ae.MakeError(ae.ERRBadRequest, errStr)
	}

	// create
	err := svc.repository.Create(song)
	if err != nil {
		log.Error("Mongo error:", err)
		return ae.MakeError(ae.ERRInternalServerError, "Could not create song")
	}

	return nil
}
