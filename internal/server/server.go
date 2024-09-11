package server

import (
	"net/http"
	"station-tracking/internal/apis"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Start() {
	r := setupServer()
	http.ListenAndServe(":8080", r)
}

func setupServer() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://frontend-134042448147.europe-west2.run.app"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))
	})

	r.Route("/stations", func(r chi.Router) {
		r.Get("/", apis.GetAllStations)
	})
	return r
}
