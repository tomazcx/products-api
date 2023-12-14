package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tomazcx/products-api/configs"
	"github.com/tomazcx/products-api/internal/infra/http/routes"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading the environment configuration: %v", err)
	}
	configs.InitializaDB()

	r := chi.NewRouter()
	router := routes.NewAppRouter(r)
	r.Use(middleware.Logger)
	router.DefineRoutes()

	fmt.Println("Server running at port 8000!")
	http.ListenAndServe(":8000", r)
}
