package utils

import (
	"math/rand"
	"time"
)

const urlAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateShortURL() string {
	const length = 8
	shortURL := make([]byte, length)
	for i := range shortURL {
		shortURL[i] = urlAlphabet[seededRand.Intn(len(urlAlphabet))]
	}
	return string(shortURL)
}
