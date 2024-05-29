package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"url-shortener/internal/cache"
	"url-shortener/internal/config"
	"url-shortener/internal/router"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize configuration
	cfg := initConfig()

	// Initialize database connection
	dbConn, cleanup := initDB(cfg.DatabaseURL)
	defer cleanup()

	// Initialize cache
	cache := cache.NewCache(5*time.Minute, 10*time.Minute)

	// Setup router
	r := initRouter(dbConn, cache, cfg.ShortURLDomains)

	// Start server
	startServer(cfg.Port, r)
}

func initConfig() *config.Config {
	config.LoadConfig()
	return &config.Config{
		Port:            config.GetEnv("PORT"),
		DatabaseURL:     config.GetEnv("DATABASE_URL"),
		ShortURLDomains: config.GetEnvAsSlice("SHORT_URL_DOMAINS", ","),
	}
}

func initDB(databaseURL string) (*sql.DB, func()) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}
	return db, func() {
		db.Close()
	}
}

func initRouter(dbConn *sql.DB, cache *cache.Cache, shortURLDomains []string) *mux.Router {
	return router.Setup(dbConn, cache, shortURLDomains)
}

func startServer(port string, r *mux.Router) {
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
