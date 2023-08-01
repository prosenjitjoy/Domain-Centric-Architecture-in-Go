package main

import (
	"context"
	"log"
	"net/http"
	"server/database"
	"server/internal/coffee"
	"server/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	ctx := context.Background()
	conn, err := database.NewDatabase(ctx)
	if err != nil {
		log.Fatal("Could not initialize database connection")
	}
	defer conn.Close(ctx)

	userAdapter := user.NewAdapter(conn.GetDB())
	userService := user.NewService(userAdapter)
	userHandler := user.NewHandler(userService)

	coffeeAdapter := coffee.NewAdapter(conn.GetDB())
	coffeeService := coffee.NewService(coffeeAdapter)
	coffeeHandler := coffee.NewHandler(coffeeService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Route("/user", func(r chi.Router) {
		r.Post("/signup", userHandler.CreateUser)
		r.Post("/signin", userHandler.Login)
		r.Get("/signout", userHandler.Logout)
	})

	router.Route("/coffee", func(r chi.Router) {
		r.Get("/", coffeeHandler.GetAllCoffees)
		r.Post("/", coffeeHandler.CreateCoffee)
		r.Get("/{uuid}", coffeeHandler.GetCoffeeByID)
		r.Put("/{uuid}", coffeeHandler.UpdateCoffeeByID)
		r.Delete("/{uuid}", coffeeHandler.DeleteCoffeeByID)
	})

	err = http.ListenAndServe(":5000", router)
	if err != nil {
		log.Fatal("Could not initialize http server")
	}
}
