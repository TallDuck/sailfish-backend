package main

import (
	"fmt"
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/tallduck/sailfish-backend/helpers"
	"github.com/tallduck/sailfish-backend/router/utils"
)

var appRouter = Router{}

func main() {
	v1 := chi.NewRouter()

	v1.Group(unauthenticatedRoutes)
	v1.Group(authenticatedRoutes)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Mount("/api/v1", v1)

	port := helpers.GetEnv("APP_PORT", "3000")
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

func unauthenticatedRoutes(r chi.Router) {
	appRouter.PublicHandlers(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("v1"))
	})
}

func authenticatedRoutes(r chi.Router) {
	r.Use(utils.AuthenticationMiddleware)
	appRouter.PrivateHandlers(r)
}
