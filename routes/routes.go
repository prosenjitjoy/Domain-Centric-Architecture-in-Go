package routes

import (
	"server/internal/coffee"
	"server/internal/user"
	"server/middlewares"

	"github.com/go-chi/chi/v5"
)

func Use(router *chi.Mux, user *user.Handler, coffee *coffee.Handler) {
	router.Group(func(r chi.Router) {
		r.Post("/signup", user.Register)
		r.Post("/signin", user.Login)
	})

	router.Group(func(r chi.Router) {
		r.Use(middlewares.Authenticator)
		r.Get("/signout", user.Logout)
		r.Get("/coffees", coffee.GetAllCoffees)
		r.Post("/coffee", coffee.CreateCoffee)
		r.Get("/coffee/{uuid}", coffee.GetCoffeeByID)
		r.Put("/coffee/{uuid}", coffee.UpdateCoffeeByID)
		r.Delete("/coffee/{uuid}", coffee.DeleteCoffeeByID)
	})
}
