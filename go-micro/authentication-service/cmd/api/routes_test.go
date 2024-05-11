package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

func Test_routes_exists(t *testing.T) {
	// Given
	testApp := Config{}

	testRoutes := testApp.routes()
	chiRoutes := testRoutes.(chi.Router)

	routes := []string{"/authenticate"}

	for _, route := range routes {
		// When
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	var found bool

	_ = chi.Walk(routes, func(method, foundRoute string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}

		return nil
	})

	if !found {
		// Then
		t.Errorf("did not find %s in registered routes", route)
	}
}
