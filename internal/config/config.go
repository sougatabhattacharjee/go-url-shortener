package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	DatabaseURL     string
	ShortURLDomains []string
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvAsSlice(key string, sep string) []string {
	value := os.Getenv(key)
	return strings.Split(value, sep)
}
