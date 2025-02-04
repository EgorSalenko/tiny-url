package main

import (
	"github.com/EgorSalenko/tiny/app"
	chi "github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	log.Info().Msg("starting server")
	a := app.New()
	r := chi.NewRouter()
	r.Post("/short", a.Routes.ShortnerHandler.GetShortUrl)
	r.Get("/*", a.Routes.ShortnerHandler.Redirect)
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
