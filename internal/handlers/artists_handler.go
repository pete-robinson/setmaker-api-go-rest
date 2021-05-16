package handlers

import (
	"net/http"
	"setmaker-api-go-rest/internal/services"

	log "github.com/sirupsen/logrus"
)

type Handler struct {
	svc *services.AppService
}

func NewArtistsHandler(svc *services.AppService) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (s *Handler) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	log.Info("HEREHERE")
}
