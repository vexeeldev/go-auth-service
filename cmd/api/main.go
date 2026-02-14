package main

import (
	"log"
	"net/http"

	"backend/internal/auth"
	"backend/internal/config"
	"backend/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()

	if err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName); err != nil {
		log.Fatal(err)
	}

	if err := database.CreateTable(); err != nil {
		log.Fatal(err)
	}

	repo := auth.NewRepository(database.DB)
	service := auth.NewService(repo)
	handler := auth.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to backend API"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Route("/api", func(r chi.Router) {
		auth.Routes(r, handler)
	})

	log.Printf("Server running at http://localhost:%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
