package repository

import (
	"database/sql"
	"errors"
	"url-shortener/internal/model"
)

type URLRepository interface {
	SaveURL(url *model.URL) error
	GetURL(shortURL string) (*model.URL, error)
	IncrementClicks(shortURL string) error
	GetAnalytics(shortURL string) (*model.Analytics, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) (*PostgresRepository, error) {
	if db == nil {
		return nil, errors.New("db connection is nil")
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) SaveURL(url *model.URL) error {
	_, err := r.db.Exec("INSERT INTO urls (short_url, long_url, created_at) VALUES ($1, $2, $3)", url.ShortURL, url.LongURL, url.CreatedAt)
	return err
}

func (r *PostgresRepository) GetURL(shortURL string) (*model.URL, error) {
	row := r.db.QueryRow("SELECT short_url, long_url, created_at FROM urls WHERE short_url = $1", shortURL)
	url := &model.URL{}
	err := row.Scan(&url.ShortURL, &url.LongURL, &url.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("URL not found")
	}
	return url, err
}

func (r *PostgresRepository) IncrementClicks(shortURL string) error {
	_, err := r.db.Exec("UPDATE urls SET clicks = clicks + 1 WHERE short_url = $1", shortURL)
	return err
}

func (r *PostgresRepository) GetAnalytics(shortURL string) (*model.Analytics, error) {
	row := r.db.QueryRow("SELECT short_url, long_url, clicks, created_at, updated_at FROM urls WHERE short_url = $1", shortURL)
	analytics := &model.Analytics{}
	err := row.Scan(&analytics.ShortURL, &analytics.LongURL, &analytics.Clicks, &analytics.CreatedAt, &analytics.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("analytics not found")
	}
	return analytics, err
}
