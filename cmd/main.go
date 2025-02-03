package main

import (
	"github.com/EgorSalenko/tiny/app"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	log.Info().Msg("starting server")
	app := app.New()
	http.HandleFunc("/", app.Routes.UrlHandler.GetUrls)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
