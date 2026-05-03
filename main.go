package main

import (
	"github.com/888tru/url-shortener/internal/handler"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Get("/ping", handler.HealthHandler)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
