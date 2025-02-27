package routes

import (
	"net/http"

	"github.com/lordaris/gotth-boilerplate/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRouter configures and returns the application router
func SetupRouter(h *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	// Static files
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Routes
	r.Get("/", h.HomeHandler())
	r.Get("/api/users", h.GetUsersHandler())
	r.Post("/api/users", h.CreateUserHandler())
	r.Get("/api/users/{id}", h.GetUserDetailHandler())
	r.Get("/api/users/{id}/edit", h.EditUserHandler())
	r.Put("/api/users/{id}", h.UpdateUserHandler())
	r.Delete("/api/users/{id}", h.DeleteUserHandler())

	return r
}
