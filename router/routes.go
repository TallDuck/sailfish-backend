package main

import (
	"net/http"

	"github.com/pressly/chi"

	"github.com/tallduck/sailfish-backend/protobuf/auth"
)

// Router is the router
type Router struct{}

// Route provides a common definintion for HTTP routes. It contains
// thier method, path, and request Protobuf
type Route struct {
	Method  string
	Path    string
	Request interface{}
}

// PublicHandlers takes a chi router and attaches the public routes
// defined inside the router.
func (r *Router) PublicHandlers(cr chi.Router) {
	r.setupHandlers(r.publicRoutes(), cr)
}

// PrivateHandlers takes a chi router and attaches the private routes
// defined inside the router.
func (r *Router) PrivateHandlers(cr chi.Router) {
	r.setupHandlers(r.privateRoutes(), cr)
}

func (r *Router) setupHandlers(routes []Route, cr chi.Router) {
	handlers := map[string]func(string, interface{}){
		"GET": func(path string, requestBuffer interface{}) {
			cr.Get(path, func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Generic GET, not implemented: " + path))
			})
		},
		"POST": func(path string, requestBuffer interface{}) {
			cr.Post(path, func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Generic POST, not implemented: " + path))
			})
		},
	}

	for _, route := range routes {
		handlers[route.Method](route.Path, route.Request)
	}
}

// publicRoutes contain route definintions for each publically
// accessibsle route that conforms to the standard hanlders
func (r *Router) publicRoutes() []Route {
	return []Route{
		Route{"GET", "/router/test", &auth.AuthRequest{}},
		Route{"POST", "/auth/login", &auth.AuthRequest{}},
	}
}

// privateRoutes contain route definintions for each route that
// requires authentication and conform to the standard hanlders
func (r *Router) privateRoutes() []Route {
	return []Route{
		Route{"GET", "/auth/logout", &auth.TokenRequest{}},
	}
}

// findByMethod takes a slice of route structs and returns the
// route definitions that match the HTTP method specified
func (r *Router) findByMethod(method string, routes []Route) []Route {
	filtered := []Route{}

	for _, item := range routes {
		if item.Method == method {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
