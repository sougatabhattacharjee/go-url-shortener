package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"url-shortener/internal/cache"
	"url-shortener/internal/config"
	"url-shortener/internal/db"
	"url-shortener/internal/router"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize configuration
	cfg := initConfig()

	// Initialize database connection
	dbConn, cleanup := db.InitDB(cfg.DatabaseURL)
	defer cleanup()

	// Initialize cache
	urlCache := cache.NewCache(cfg.CacheExpiration, 10*time.Minute)

	// Setup router
	r := initRouter(dbConn, urlCache, cfg.ShortURLDomains, cfg.CacheExpiration)

	// Start server
	startServer(cfg.Port, r)
}

func initConfig() *config.Config {
	config.LoadConfig()
	cfg := &config.Config{
		Port:            config.GetEnv("PORT"),
		DatabaseURL:     config.GetEnv("DATABASE_URL"),
		ShortURLDomains: config.GetEnvAsSlice("SHORT_URL_DOMAINS", ","),
		CacheExpiration: config.GetEnvAsDuration("CACHE_EXPIRATION"),
	}
	log.Printf("Loaded configuration: %+v", cfg) // Add logging
	return cfg
}

func initRouter(dbConn *sql.DB, urlCache *cache.Cache, shortURLDomains []string, cacheExpiration time.Duration) *mux.Router {
	return router.Setup(dbConn, urlCache, shortURLDomains, cacheExpiration)
}

func startServer(port string, r *mux.Router) {
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
