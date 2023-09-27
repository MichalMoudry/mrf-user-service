package transport

import (
	"net/http"
	"user-service/internal/transport/model"

	"firebase.google.com/go/v4/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Port     int
	Mux      *chi.Mux
	Services *model.ServiceCollection
}

// Initializer function for HTTP handler.
func Initalize(port int, services *model.ServiceCollection, auth *auth.Client) *Handler {
	handler := &Handler{
		Port:     port,
		Mux:      chi.NewRouter(),
		Services: services,
	}
	handler.Mux.Use(middleware.Logger)

	// Protected routes
	handler.Mux.Group(func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Delete("/", handler.DeleteUser)
		})
	})

	// Public routes
	handler.Mux.Get("/health", health)
	return handler
}

// Controller endpoint function for handling requests on /health.
func health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNoContent)
}
