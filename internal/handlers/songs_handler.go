package handlers

import (
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

type SongsHandler struct {
	svc *services.SongService
}

type SongsList []*domain.Song

func NewSongsHandler(svc *services.SongService) *SongsHandler {
	return &SongsHandler{
		svc: svc,
	}
}

func (s *SongsHandler) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	route := mux.CurrentRoute(r).GetName()
	var successCode int
	var resp interface{}
	var err *ae.AppError

	switch route {
	case "createSong":
		resp, err = s.createSong(r)
		successCode = 201
		break
	case "getSong":
		resp, err = s.getSong(r)
		successCode = 200
		break
	}

	if err != nil {
		utils.JsonResponse(w, err.GetCode(), err)
	} else {
		utils.JsonResponse(w, successCode, resp)
	}
}

func (s *SongsHandler) getSong(r *http.Request) (interface{}, *ae.AppError) {
	idb := mux.Vars(r)["id"]
	id, e := uuid.Parse(idb)
	if e != nil {
		log.Error(e)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	var song *domain.Song
	song, err := s.svc.GetSong(id)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *SongsHandler) createSong(r *http.Request) (interface{}, *ae.AppError) {
	var song *domain.Song

	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		log.Error(err)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid request body")
	}

	if err := s.svc.CreateSong(song); err != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, err)
	}

	return song, nil
}
