package model

type ShortenRequest struct {
	LongURL     string `json:"long_url"`
	CustomAlias string `json:"custom_alias,omitempty"`
	Domain      string `json:"domain"`
}

type URL struct {
	ShortURL         string `json:"short_url"`
	LongURL          string `json:"long_url"`
	Domain           string `json:"domain"`
	CompleteShortURL string `json:"complete_short_url,omitempty"`
	Clicks           int    `json:"clicks"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type Analytics struct {
	ShortURL         string `json:"short_url"`
	LongURL          string `json:"long_url"`
	Domain           string `json:"domain"`
	CompleteShortURL string `json:"complete_short_url,omitempty"`
	Clicks           int    `json:"clicks"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
