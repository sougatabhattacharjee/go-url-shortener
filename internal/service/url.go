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

	// Check if the combination of longURL and domain already exists
	existingURL, err := s.repo.GetURLByLongURLAndDomain(longURL, domain)
	if err == nil {
		return fmt.Sprintf("%s/%s", existingURL.Domain, existingURL.ShortURL), nil
	}

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
