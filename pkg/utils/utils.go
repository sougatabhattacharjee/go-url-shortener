package utils

import (
	"math/rand"
	"time"
)

const urlAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateShortURL() string {
	const length = 8
	shortURL := make([]byte, length)
	for i := range shortURL {
		shortURL[i] = urlAlphabet[rand.Intn(len(urlAlphabet))]
	}
	return string(shortURL)
}
