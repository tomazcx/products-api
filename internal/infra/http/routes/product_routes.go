package routes

import (
	"github.com/go-chi/chi"
	"github.com/tomazcx/products-api/internal/data/factory"
	"github.com/tomazcx/products-api/internal/infra/http/handlers"
)

func UseProductRoutes(r *chi.Mux) {
	factory := factory.ProductFactory{}
	productHandler := handlers.NewProductHandler(factory)

	r.Get("/products", productHandler.GetManyProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
}
