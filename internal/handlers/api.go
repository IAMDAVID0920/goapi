package handlers

import (
	"goapi/internal/middleware"

	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

// function with Capitalized name means we can import in other places as packages.
// lowercase such as handler() {} will be a private function
func Handler(r *chi.Mux) {
	// Global middleware
	// this is use to strip the trailing slashes in case
	// we type https://localhost:8000/account/coins`/` (404) error code
	// instead of https://localhost:8000/account/coins (200) success code
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance) // define this http handler function later
	})
}
