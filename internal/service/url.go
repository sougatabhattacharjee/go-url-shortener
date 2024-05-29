package service

import (
	"errors"
	"fmt"
	"time"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	"url-shortener/pkg/utils"
)

type URLService struct {
	repo            repository.URLRepository
	shortURLDomains []string
}

func NewURLService(repo repository.URLRepository, shortURLDomains []string) *URLService {
	return &URLService{
		repo:            repo,
		shortURLDomains: shortURLDomains,
	}
}

func (s *URLService) ShortenURL(longURL, customAlias, domain string) (string, error) {
	if !s.isValidDomain(domain) {
		return "", errors.New("unsupported domain")
	}

	shortURL := customAlias
	if shortURL == "" {
		shortURL = utils.GenerateShortURL()
	}

	fullShortURL := fmt.Sprintf("%s/%s", domain, shortURL)

	url := &model.URL{
		ShortURL:  shortURL,
		LongURL:   longURL,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err := s.repo.SaveURL(url)
	if err != nil {
		return "", err
	}

	return fullShortURL, nil
}

func (s *URLService) isValidDomain(domain string) bool {
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

	s.repo.IncrementClicks(shortURL)

	return url.LongURL, nil
}

func (s *URLService) GetURLDetails(shortURL string) (*model.URL, error) {
	return s.repo.GetURL(shortURL)
}

func (s *URLService) GetAnalytics(shortURL string) (*model.Analytics, error) {
	return s.repo.GetAnalytics(shortURL)
}
