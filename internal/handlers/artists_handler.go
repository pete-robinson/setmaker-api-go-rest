package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/services"
	"setmaker-api-go-rest/internal/utils"
	e "setmaker-api-go-rest/internal/utils/errors"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Handler struct {
	svc *services.ArtistService
}

type ArtistsList []*domain.Artist

func NewArtistsHandler(svc *services.ArtistService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (s *Handler) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	route := mux.CurrentRoute(r).GetName()
	var successCode int
	var resp interface{}
	var err e.AppError

	switch route {
	case "createArtist":
		resp, err = s.createArtist(r)
		successCode = 201
		break
	case "getArtists":
		resp, err = s.getArtists(r)
		successCode = 200
		break
	case "getArtist":
		resp, err = s.getArtist(r)
		successCode = 200
		break
	default:
		fmt.Println("no")
		break
	}

	if err != nil {
		utils.JsonResponse(w, err.GetCode(), err.GetPayload())
	} else {
		utils.JsonResponse(w, successCode, resp)
	}
}

func (s *Handler) getArtists(r *http.Request) (interface{}, e.AppError) {
	sort := utils.FetchSortParams(r, "name", 1)

	artists, err := s.svc.GetArtists(sort)
	if err != nil {
		log.Error(err)
		return nil, e.NewInternalServerError("An internal error occurred", err)
	}

	return artists, nil
}

func (s *Handler) createArtist(r *http.Request) (interface{}, e.AppError) {
	var a *domain.Artist

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Error(err)
		return nil, e.NewInternalServerError("An internal error occurred", err)
	}

	if err := s.svc.CreateArtist(a); err != nil {
		return nil, e.NewBadRequest("Bad Request", err)
	}

	return a, nil
}

func (s *Handler) getArtist(r *http.Request) (interface{}, e.AppError) {
	idb := mux.Vars(r)["id"]
	id, err := uuid.Parse(idb)
	if err != nil {
		log.Error(err)
		return nil, e.NewBadRequest("Bad Request", "Invalid ID provided")
	}

	var artist *domain.Artist
	if artist, err = s.svc.GetArtist(id); err != nil {
		return nil, e.NewNotFound(fmt.Sprintf("Artist %s not found", id), err)
	}

	return artist, nil
}
