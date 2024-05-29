package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

func Setup(db *sql.DB, shortURLDomains []string) *mux.Router {
	repo, err := repository.NewPostgresRepository(db)
	if err != nil {
		log.Fatalf("Could not create repository: %v", err)
	}
	urlService := service.NewURLService(repo, shortURLDomains)
	urlHandler := handler.NewURLHandler(urlService)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/shorten", urlHandler.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", urlHandler.Redirect).Methods("GET")
	r.HandleFunc("/api/v1/urls/{shortURL}", urlHandler.GetURLDetails).Methods("GET")
	r.HandleFunc("/api/v1/urls/{shortURL}/analytics", urlHandler.GetAnalytics).Methods("GET")

	return r
}
