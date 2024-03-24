package net

import (
	"fmt"
	"net/http"
	"network-programming/net/handlers"

	"github.com/go-chi/chi"

	"github.com/go-chi/cors"
)

func Run(port uint) {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	r.Route("/", func(r chi.Router) {
		r.Post("/sort_words", handlers.ReplaceRoundBrackets)
	})

	startHttpListener(port, r)
}

func startHttpListener(port uint, r *chi.Mux) {
	fmt.Printf("Server is listening on port %d...\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
