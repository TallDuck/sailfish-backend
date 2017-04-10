package main

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/tallduck/sailfish-backend/router/utils"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))

	r.Group(unauthenticatedRoutes)
	r.Group(authenticatedRoutes)

	http.ListenAndServe(":3000", r)
}

func unauthenticatedRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("v1"))
	})

	r.Get("/auth/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login here"))
	})
}

func authenticatedRoutes(r chi.Router) {
	r.Use(utils.AuthenticationMiddleware)

	r.Get("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("logout here"))
	})
}
