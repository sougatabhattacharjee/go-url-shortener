package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"url-shortener/internal/service"
	"url-shortener/pkg/utils"
)

type URLHandler struct {
	service *service.URLService
}

func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req service.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.ShortenURL(req.LongURL, req.CustomAlias, req.Domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	longURL, err := h.service.GetLongURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func (h *URLHandler) GetURLDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	urlDetails, err := h.service.GetURLDetails(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(urlDetails)
}

func (h *URLHandler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	analytics, err := h.service.GetAnalytics(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(analytics)
}

func (h *URLHandler) GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	urlDetails, err := h.service.GetURLDetails(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	qrCode, err := utils.GenerateQRCode(urlDetails.CompleteShortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(qrCode)
}
