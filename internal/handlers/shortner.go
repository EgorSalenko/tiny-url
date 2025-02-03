package handlers

import (
	"encoding/json"
	"github.com/EgorSalenko/tiny/internal/shortener"
	"github.com/rs/zerolog/log"
	"net/http"
)

type UrlHandler struct {
	service *shortener.Service
}

func NewUrlHandler(s *shortener.Service) *UrlHandler {
	return &UrlHandler{service: s}
}

func (h *UrlHandler) GetUrls(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error().Err(err).Msg("error parsing form")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	originalUrl := r.Form.Get("url")
	if originalUrl == "" {
		log.Error().Msg("Empty form")
		http.Error(w, "empty url", http.StatusBadRequest)
		return
	}
	hash, err := h.service.GetUrlHash(r.Context(), originalUrl)
	if err != nil {
		log.Error().Err(err).Msg("error getting hash")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(map[string]string{originalUrl: hash})
	if err != nil {
		log.Error().Err(err).Msg("error encoding url")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
