package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"setmaker-api-go-rest/internal/domain"
	"setmaker-api-go-rest/internal/services"
	"setmaker-api-go-rest/internal/utils"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Handler struct {
	svc *services.ArtistService
}

func NewArtistsHandler(svc *services.ArtistService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (s *Handler) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	route := mux.CurrentRoute(r).GetName()
	var payload interface{}
	var code int

	switch route {
	case "createArtist":
		payload, code = s.createArtist(r)
		break
	default:
		fmt.Println("no")
		break
	}

	utils.JsonResponse(w, code, payload)
}

func (s *Handler) createArtist(r *http.Request) (interface{}, int) {
	var a *domain.Artist
	var code int
	var resp interface{}

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.Error(err)
		code = 500
		resp = err
	}

	if ok, err := s.svc.CreateArtist(a); ok {
		code = 201
		resp = a
	} else {
		code = 400
		resp = utils.CreateErrorPayload(err)
	}

	return resp, code
}
