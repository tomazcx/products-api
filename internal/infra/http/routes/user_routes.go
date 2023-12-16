package routes

import (
	"github.com/go-chi/chi"
	"github.com/tomazcx/products-api/internal/data/factory"
	"github.com/tomazcx/products-api/internal/infra/http/handlers"
)

func UseUserRoutes(r *chi.Mux) {
	factory := factory.UserFactory{}
	userHandler := handlers.NewUserHandler(factory)

	r.Get("/users/{email}", userHandler.GetByEmail)
	r.Post("/users", userHandler.Create)
	r.Post("/users/login", userHandler.Authenticate)
}
