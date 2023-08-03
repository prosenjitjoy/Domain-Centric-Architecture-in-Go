package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/database"
	"server/internal/coffee"
	"server/internal/user"
	"server/middlewares"
	"server/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	router.Use(middlewares.Cors)
	routes.Use(router, userHandler, coffeeHandler)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("Could not initialize http server")
	}
}
