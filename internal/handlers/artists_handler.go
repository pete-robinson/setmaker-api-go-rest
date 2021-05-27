package handlers

import (
	"encoding/json"
	"net/http"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/services"
	"setmaker-api-go-rest/internal/utils"
	ae "setmaker-api-go-rest/internal/utils/errors"

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
	route := mux.CurrentRoute(r).GetName() // current requested route
	var successCode int                    // init success code
	var resp interface{}                   // init response interface
	var err *ae.AppError                   // init pointer to error

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
	case "updateArtist":
		resp, err = s.updateArtist(r)
		successCode = 200
	case "deleteArtist":
		resp, err = s.deleteArtist(r)
		successCode = 200
	default:
		err = ae.MakeError(ae.ERRNotFound, "Route not found")
		break
	}

	// @todo can probably combine this and rely on interfaces to do the work
	if err != nil {
		utils.JsonResponse(w, err.GetCode(), err)
	} else {
		utils.JsonResponse(w, successCode, resp)
	}
}

func (s *Handler) getArtists(r *http.Request) (interface{}, *ae.AppError) {
	// fetch and parse sort params
	sort, e := utils.FetchSortParams(r, "name", 1)
	if e != nil {
		// throw error if invalid format
		return nil, ae.MakeError(ae.ERRBadRequest, e)
	}

	artists, err := s.svc.GetArtists(sort)
	if err != nil {
		log.Error(err)
		return nil, ae.MakeError(ae.ERRBadRequest, err)
	}

	return artists, nil
}

func (s *Handler) getArtist(r *http.Request) (interface{}, *ae.AppError) {
	// fetch and parse ID from URL
	idb := mux.Vars(r)["id"]
	id, e := uuid.Parse(idb)
	if e != nil {
		log.Error(e)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	// init artist pointer
	var artist *domain.Artist
	artist, err := s.svc.GetArtist(id)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (s *Handler) createArtist(r *http.Request) (interface{}, *ae.AppError) {
	var a *domain.Artist

	// attempt to unmarshal request body into artist struct
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Error(err)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid request body")
	}

	if err := s.svc.CreateArtist(a); err != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, err)
	}

	return a, nil
}

func (s *Handler) updateArtist(r *http.Request) (interface{}, *ae.AppError) {
	// fetch and parse id from url
	var a *domain.Artist
	idb := mux.Vars(r)["id"]
	id, err := uuid.Parse(idb)
	if err != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	// attempt to unmarshal request body into artist struct
	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Error(err)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid request body")
	}

	if err := s.svc.UpdateArtist(a, id); err != nil {
		return nil, err
	}

	return a, nil
}

func (s *Handler) deleteArtist(r *http.Request) (interface{}, *ae.AppError) {
	// fetch and parse id from url
	idb := mux.Vars(r)["id"]
	id, e := uuid.Parse(idb)
	if e != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	artist, err := s.svc.DeleteArtist(id)
	if err != nil {
		return nil, err
	}

	return artist, nil

}
