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
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type artistsHandler struct {
	svc services.ArtistService
}

type ArtistsList []*domain.Artist

const (
	routeCreateArtist = "createArtist"
	routeGetArtists   = "getArtists"
	routeGetArtist    = "getArtist"
	routeUpdateArtist = "updateArtist"
	routeDeleteArtist = "deleteArtist"
)

/**
 * Instantiate new service
 */
func NewArtistsHandler(svc services.ArtistService) *artistsHandler {
	return &artistsHandler{
		svc: svc,
	}
}

/**
 * Handle inbound HTTP routes
 */
func (s *artistsHandler) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	route := mux.CurrentRoute(r).GetName() // current requested route
	var successCode int                    // init success code
	var resp interface{}                   // init response interface
	var err *ae.AppError                   // init error

	switch route {
	case routeCreateArtist:
		resp, err = s.createArtist(ctx, r)
		successCode = 201
		break
	case routeGetArtists:
		resp, err = s.getArtists(ctx, r)
		successCode = 200
		break
	case routeGetArtist:
		resp, err = s.getArtist(ctx, r)
		successCode = 200
		break
	case routeUpdateArtist:
		resp, err = s.updateArtist(ctx, r)
		successCode = 200
	case routeDeleteArtist:
		resp, err = s.deleteArtist(ctx, r)
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

/**
 * get a list of artists
 */
func (s *artistsHandler) getArtists(ctx context.Context, r *http.Request) ([]*domain.Artist, *ae.AppError) {
	// fetch and parse sort params
	sort, e := utils.FetchSortParams(r, "name", 1)
	if e != nil {
		// throw error if invalid format
		return nil, ae.MakeError(ae.ERRBadRequest, e)
	}

	artists, err := s.svc.GetArtists(ctx, sort)
	if err != nil {
		log.Error(err)
		return nil, ae.MakeError(ae.ERRBadRequest, err)
	}

	return artists, nil
}

/**
 * Get single artist by ID
 */
func (s *artistsHandler) getArtist(ctx context.Context, r *http.Request) (*domain.Artist, *ae.AppError) {
	// fetch and parse ID from URL
	idb := mux.Vars(r)["id"]
	id, e := uuid.Parse(idb)
	if e != nil {
		log.Error(e)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	// init artist pointer
	var artist *domain.Artist
	artist, err := s.svc.GetArtist(ctx, id)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

/**
 * Create a new artist
 */
func (s *artistsHandler) createArtist(ctx context.Context, r *http.Request) (*domain.Artist, *ae.AppError) {
	var a *domain.Artist

	// attempt to unmarshal request body into artist struct
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Error(err)
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid request body")
	}

	if err := s.svc.CreateArtist(ctx, a); err != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, err)
	}

	return a, nil
}

/**
 * Update existing artist
 */
func (s *artistsHandler) updateArtist(ctx context.Context, r *http.Request) (*domain.Artist, *ae.AppError) {
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

	if err := s.svc.UpdateArtist(ctx, a, id); err != nil {
		return nil, err
	}

	return a, nil
}

/**
 * Delete an artist
 */
func (s *artistsHandler) deleteArtist(ctx context.Context, r *http.Request) (*domain.Artist, *ae.AppError) {
	// fetch and parse id from url
	idb := mux.Vars(r)["id"]
	id, e := uuid.Parse(idb)
	if e != nil {
		return nil, ae.MakeError(ae.ERRBadRequest, "Invalid ID")
	}

	artist, err := s.svc.DeleteArtist(ctx, id)
	if err != nil {
		return nil, err
	}

	return artist, nil

}
