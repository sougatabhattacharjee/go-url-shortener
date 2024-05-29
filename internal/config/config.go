package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	DatabaseURL     string
	ShortURLDomains []string
	CacheExpiration time.Duration
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	logEnvVariables()
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvAsSlice(key string, sep string) []string {
	value := os.Getenv(key)
	if value == "" {
		return []string{}
	}
	return strings.Split(value, sep)
}

func GetEnvAsDuration(key string) time.Duration {
	value := os.Getenv(key)
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 5 * time.Minute // default value
	}
	return duration
}

func logEnvVariables() {
	log.Printf("PORT: %s", GetEnv("PORT"))
	log.Printf("DATABASE_URL: %s", GetEnv("DATABASE_URL"))
	log.Printf("SHORT_URL_DOMAINS: %s", GetEnv("SHORT_URL_DOMAINS"))
	log.Printf("CACHE_EXPIRATION: %s", GetEnv("CACHE_EXPIRATION"))
}
