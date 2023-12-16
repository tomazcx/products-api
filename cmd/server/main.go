package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/tomazcx/products-api/configs"
	_ "github.com/tomazcx/products-api/docs"
	"github.com/tomazcx/products-api/internal/infra/http/routes"
)

// @title Products API
// @version 1.0
// @description CRUD of products with user authentication.

// @contact.name Tomaz C. Xavier
// @contact.url http://www.tomazcx.site
// @contact.email tomazcx06@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading the environment configuration: %v", err)
	}
	configs.InitializaDB()

	r := chi.NewRouter()
	router := routes.NewAppRouter(r)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenAuth))
	r.Use(middleware.WithValue("exp", config.JWTExpiresIn))
	router.DefineRoutes()

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/doc.json"),
	))
	fmt.Println("Server running at port 8000!")
	http.ListenAndServe(":8000", r)
}
