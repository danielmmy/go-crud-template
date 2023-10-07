package handlers

import (
	"net/http"

	"github.com/danielmmy/templates/gocrud/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Handle is the app handler.
// it returns an http handler.
func (app *App) Handle() http.Handler {
	mux := chi.NewMux()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(chimiddle.StripSlashes)
	mux.Use(chimiddle.Heartbeat("/health-check"))
	mux.Route("/users", func(route chi.Router) {
		route.Use(middleware.Authorize)
		route.Post("/", app.addUser)
		route.Get("/", app.listUsers)
		route.Get("/{username}", app.getUser)
		route.Put("/{username}", app.updateUser)
		route.Delete("/{username}", app.deleteUser)
	})

	return mux
}
