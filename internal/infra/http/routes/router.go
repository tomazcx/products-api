package routes

import "github.com/go-chi/chi"

type AppRouter struct {
	r *chi.Mux
}

func NewAppRouter(r *chi.Mux) *AppRouter {
	return &AppRouter{r: r}
}

func (a *AppRouter) DefineRoutes() {
	UseProductRoutes(a.r)
	UseUserRoutes(a.r)
}
