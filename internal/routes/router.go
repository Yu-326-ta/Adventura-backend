package routes

import (
	handler "Adventura/internal/handlers/health"
	"Adventura/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(hu usecase.Health) http.Handler {
	r := chi.NewRouter()

	r.Mount("/api/health", handler.NewRouter(hu))

	return r
}
