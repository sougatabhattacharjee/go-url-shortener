package repository

import (
	"database/sql"
	"url-shortener/internal/model"
)

type URLRepository interface {
	SaveURL(url *model.URL) error
	GetURL(shortURL string) (*model.URL, error)
	GetURLByLongURLAndDomain(longURL, domain string) (*model.URL, error) // New method
	IncrementClicks(shortURL string) error
	GetAnalytics(shortURL string) (*model.Analytics, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) (URLRepository, error) {
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) SaveURL(url *model.URL) error {
	query := `
        INSERT INTO urls (short_url, long_url, domain, created_at)
        VALUES ($1, $2, $3, $4)
    `
	_, err := r.db.Exec(query, url.ShortURL, url.LongURL, url.Domain, url.CreatedAt)
	return err
}

func (r *PostgresRepository) GetURL(shortURL string) (*model.URL, error) {
	query := `
        SELECT short_url, long_url, domain, clicks, created_at, updated_at
        FROM urls
        WHERE short_url = $1
    `
	row := r.db.QueryRow(query, shortURL)
	url := &model.URL{}
	err := row.Scan(&url.ShortURL, &url.LongURL, &url.Domain, &url.Clicks, &url.CreatedAt, &url.UpdatedAt)
	return url, err
}

func (r *PostgresRepository) GetURLByLongURLAndDomain(longURL, domain string) (*model.URL, error) {
	query := `
        SELECT short_url, long_url, domain, clicks, created_at, updated_at
        FROM urls
        WHERE long_url = $1 AND domain = $2
    `
	row := r.db.QueryRow(query, longURL, domain)
	url := &model.URL{}
	err := row.Scan(&url.ShortURL, &url.LongURL, &url.Domain, &url.Clicks, &url.CreatedAt, &url.UpdatedAt)
	return url, err
}

func (r *PostgresRepository) IncrementClicks(shortURL string) error {
	query := `
        UPDATE urls
        SET clicks = clicks + 1, updated_at = NOW()
        WHERE short_url = $1
    `
	_, err := r.db.Exec(query, shortURL)
	return err
}

func (r *PostgresRepository) GetAnalytics(shortURL string) (*model.Analytics, error) {
	query := `
        SELECT short_url, long_url, domain, clicks, created_at, updated_at
        FROM urls
        WHERE short_url = $1
    `
	row := r.db.QueryRow(query, shortURL)
	analytics := &model.Analytics{}
	err := row.Scan(&analytics.ShortURL, &analytics.LongURL, &analytics.Domain, &analytics.Clicks, &analytics.CreatedAt, &analytics.UpdatedAt)
	return analytics, err
}
