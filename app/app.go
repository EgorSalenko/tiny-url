package app

import (
	"github.com/EgorSalenko/tiny/internal/handlers"
	"github.com/EgorSalenko/tiny/internal/shortener"
	"github.com/EgorSalenko/tiny/storage"
	"github.com/rs/zerolog/log"
)

type App struct {
	Routes *routes
}
type routes struct {
	ShortnerHandler *handlers.UrlHandler
}

func New() *App {
	redis := storage.NewStorage()
	_, err := redis.Ping()
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to redis")
		return nil
	}
	return &App{
		Routes: initRoutes(initServices(redis)),
	}
}

type services struct {
	shortener *shortener.Service
}

func initServices(storage *storage.Storage) services {
	return services{
		shortener: shortener.NewService(storage),
	}
}

func initRoutes(s services) *routes {
	shortnerHandler := handlers.NewShortnerHandler(s.shortener)
	return &routes{ShortnerHandler: shortnerHandler}
}
