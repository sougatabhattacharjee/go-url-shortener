package service

import (
	"errors"
	"fmt"
	"log"
	"time"
	"url-shortener/internal/cache"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	"url-shortener/pkg/utils"
)

type URLService struct {
	repo            repository.URLRepository
	cache           *cache.Cache
	shortURLDomains []string
	cacheExpiration time.Duration
}

func NewURLService(repo repository.URLRepository, cache *cache.Cache, shortURLDomains []string, cacheExpiration time.Duration) *URLService {
	log.Printf("Initializing URLService with domains: %v", shortURLDomains) // Add logging
	return &URLService{
		repo:            repo,
		cache:           cache,
		shortURLDomains: shortURLDomains,
		cacheExpiration: cacheExpiration,
	}
}

func (s *URLService) ShortenURL(longURL, customAlias, domain string) (string, error) {
	if !s.isValidDomain(domain) {
		return "", errors.New("unsupported domain")
	}

	cacheKey := fmt.Sprintf("%s:%s", domain, longURL)
	if cachedShortURL, found := s.cache.Get(cacheKey); found {
		log.Printf("Cache hit: %s", cacheKey)
		return cachedShortURL, nil
	}
	log.Printf("Cache miss: %s", cacheKey)

	// Check if the combination of longURL and domain already exists
	existingURL, err := s.repo.GetURLByLongURLAndDomain(longURL, domain)
	if err == nil {
		shortURL := fmt.Sprintf("%s/%s", existingURL.Domain, existingURL.ShortURL)
		s.cache.Set(cacheKey, shortURL, s.cacheExpiration)
		log.Printf("Database hit: %s", cacheKey)
		return shortURL, nil
	}
	log.Printf("Database miss: %s", cacheKey)

	shortURL := customAlias
	if shortURL == "" {
		shortURL = utils.GenerateShortURL()
	}

	fullShortURL := fmt.Sprintf("%s/%s", domain, shortURL)

	url := &model.URL{
		ShortURL:  shortURL,
		LongURL:   longURL,
		Domain:    domain,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err = s.repo.SaveURL(url)
	if err != nil {
		return "", err
	}

	s.cache.Set(cacheKey, fullShortURL, s.cacheExpiration)
	log.Printf("Generated and cached: %s", cacheKey)

	return fullShortURL, nil
}

func (s *URLService) isValidDomain(domain string) bool {
	log.Printf("Validating domain: %s", domain) // Add logging
	for _, d := range s.shortURLDomains {
		if d == domain {
			return true
		}
	}
	return false
}

func (s *URLService) GetLongURL(shortURL string) (string, error) {
	url, err := s.repo.GetURL(shortURL)
	if err != nil {
		return "", err
	}

	if err := s.repo.IncrementClicks(shortURL); err != nil {
		log.Printf("Failed to increment clicks for %s: %v", shortURL, err)
	}

	return url.LongURL, nil
}

func (s *URLService) GetURLDetails(shortURL string) (*model.URL, error) {
	url, err := s.repo.GetURL(shortURL)
	if err != nil {
		return nil, err
	}
	url.CompleteShortURL = fmt.Sprintf("%s/%s", url.Domain, url.ShortURL)
	return url, nil
}

func (s *URLService) GetAnalytics(shortURL string) (*model.Analytics, error) {
	analytics, err := s.repo.GetAnalytics(shortURL)
	if err != nil {
		return nil, err
	}
	analytics.CompleteShortURL = fmt.Sprintf("%s/%s", analytics.Domain, analytics.ShortURL)
	return analytics, nil
}
