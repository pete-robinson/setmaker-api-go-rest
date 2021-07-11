package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/services"
	"setmaker-api-go-rest/internal/utils"
	ae "setmaker-api-go-rest/internal/utils/errors"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type songsHandler struct {
	svc services.SongService
}

type SongsList []*domain.Song

func NewSongsHandler(svc services.SongService) *songsHandler {
	return &songsHandler{
		svc: svc,
	}
}

const (
	routeCreateSong       = "createSong"
	routeGetSong          = "getSong"
	routeUpdateSong       = "updateSong"
	routeDeleteSong       = "deleteSong"
	routeGetSongsByArtist = "getSongsByArtist"
)

func (s *songsHandler) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	route := mux.CurrentRoute(r).GetName()
	var successCode int
	var resp interface{}
	var err *ae.AppError

	switch route {
	case routeGetSongsByArtist:
		resp, err = s.getSongsByArtist(ctx, r)
		successCode = 200
		break
	case routeCreateSong:
		resp, err = s.createSong(ctx, r)
		successCode = 201
		break
	case routeGetSong:
		resp, err = s.getSong(ctx, r)
		successCode = 200
		break
	case routeUpdateSong:
		resp, err = s.updateSong(ctx, r)
		successCode = 200
		break
	case routeDeleteSong:
		resp, err = s.deleteSong(ctx, r)
		successCode = 200
		break
	}

	if err != nil {
		utils.JsonResponse(w, err.GetCode(), err)
	} else {
		utils.JsonResponse(w, successCode, resp)
	}
}

func (s *songsHandler) getSongsByArtist(ctx context.Context, r *http.Request) ([]*domain.Song, *ae.AppError) {
	idb := mux.Vars(r)["id"]
	id, e := uuid.Parse(idb)
	if e != nil {
		log.Error(e)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid artist ID")
	}

	var results []*domain.Song
	results, err := s.svc.GetSongsByArtistId(ctx, id)
	if err != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid artist")
	}

	return results, nil
}

func (s *songsHandler) getSong(ctx context.Context, r *http.Request) (*domain.Song, *ae.AppError) {
	idb := mux.Vars(r)["id"]
	id, e := uuid.Parse(idb)
	if e != nil {
		log.Error(e)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	var song *domain.Song
	song, err := s.svc.GetSong(ctx, id)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *songsHandler) createSong(ctx context.Context, r *http.Request) (*domain.Song, *ae.AppError) {
	var song *domain.Song

	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		log.Error(err)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid request body")
	}

	if err := s.svc.CreateSong(ctx, song); err != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, err)
	}

	return song, nil
}

func (s *songsHandler) updateSong(ctx context.Context, r *http.Request) (*domain.Song, *ae.AppError) {
	// fetch and parse ID from URL
	var song *domain.Song
	idb := mux.Vars(r)["id"]
	id, err := uuid.Parse(idb)
	if err != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	// unmarshal request body
	err = json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		log.Error(err)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid request body")
	}

	if err := s.svc.UpdateSong(ctx, song, id); err != nil {
		return nil, err
	}

	return song, nil
}

func (s *songsHandler) deleteSong(ctx context.Context, r *http.Request) (*domain.Song, *ae.AppError) {
	// fetch and parse ID from URL
	idb := mux.Vars(r)["id"]
	id, e := uuid.Parse(idb)
	if e != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	song, err := s.svc.DeleteSong(ctx, id)
	if err != nil {
		return nil, err
	}

	return song, nil
}
