package main

import (
	"context"
	"github.com/888tru/url-shortener/internal/handler"
	"github.com/888tru/url-shortener/internal/repository"
	"github.com/888tru/url-shortener/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	ctx := context.Background()

	db, err := pgxpool.New(ctx, "postgres://admin:secret@localhost:5432/urlshortener")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewPostgresRepository(db)
	_ = repo

	svc := service.NewUrlService(repo)
	_ = svc
	router.Get("/ping", handler.HealthHandler)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
