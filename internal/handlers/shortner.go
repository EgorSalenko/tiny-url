package handlers

import (
	"encoding/json"
	"github.com/EgorSalenko/tiny/internal/shortener"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

type UrlHandler struct {
	service *shortener.Service
}

func NewShortnerHandler(s *shortener.Service) *UrlHandler {
	return &UrlHandler{service: s}
}

func (h *UrlHandler) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("Getting short url")
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
	data, err := h.service.Hash(r.Context(), originalUrl)
	if err != nil {
		log.Error().Err(err).Interface("data", data).Msg("error getting hash")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error().Err(err).Msg("error encoding url")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Info().Interface("data", data).Msgf("Successfully got short url: %s", originalUrl)
}

func (h *UrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	hash := strings.TrimPrefix(path, "/")
	log.Info().Msgf("Redirecting to %s", hash)

	url, err := h.service.GetUrlByHash(r.Context(), hash)
	if err != nil {
		log.
			Error().
			Err(err).
			Msg("Original url by this hash does not exist")
		http.Error(w, "Original url by this hash does not exist", http.StatusNotFound)
		return
	}

	log.Info().
		Str("path", path).
		Str("url", url).
		Msg("Redirecting to url")

	http.Redirect(w, r, url, http.StatusFound)
}
