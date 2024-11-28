package handlers

import (
	"Adventura/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	healthUsecase usecase.Health
}

func NewRouter(hu usecase.Health) http.Handler {
	r := chi.NewRouter()
	h := &handler{
		healthUsecase: hu,
	}
	r.Get("/", h.Health)

	return r
}
