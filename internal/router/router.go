package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"time"
	"url-shortener/internal/cache"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

func Setup(db *sql.DB, cache *cache.Cache, shortURLDomains []string, cacheExpiration time.Duration) *mux.Router {
	repo, err := repository.NewPostgresRepository(db)
	if err != nil {
		log.Fatalf("Could not create repository: %v", err)
	}
	urlService := service.NewURLService(repo, cache, shortURLDomains, cacheExpiration)
	urlHandler := handler.NewURLHandler(urlService)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/shorten", urlHandler.ShortenURL).Methods("POST")
	r.HandleFunc("/{shortURL}", urlHandler.Redirect).Methods("GET")
	r.HandleFunc("/api/v1/urls/{shortURL}", urlHandler.GetURLDetails).Methods("GET")
	r.HandleFunc("/api/v1/urls/{shortURL}/analytics", urlHandler.GetAnalytics).Methods("GET")
	r.HandleFunc("/api/v1/urls/{shortURL}/qrcode", urlHandler.GenerateQRCode).Methods("GET")

	return r
}
