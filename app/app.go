package app

import (
	"github.com/EgorSalenko/tiny/internal/handlers"
	"github.com/EgorSalenko/tiny/internal/shortener"
	"github.com/EgorSalenko/tiny/internal/storage"
)

type App struct {
	Routes *routes
}
type routes struct {
	UrlHandler *handlers.UrlHandler
}

func New() *App {
	redis := storage.NewStorage()
	return &App{
		Routes: initHandlers(initServices(initRepositories(redis))),
	}
}

type services struct {
	shortener *shortener.Service
}

type repositories struct {
	shortener shortener.Repository
}

func initRepositories(storage *storage.Storage) *repositories {
	return &repositories{
		shortener: shortener.NewRepository(storage),
	}
}

func initServices(r *repositories) services {
	return services{
		shortener: shortener.NewService(r.shortener),
	}
}

func initHandlers(s services) *routes {
	urlHandler := handlers.NewUrlHandler(s.shortener)
	return &routes{
		UrlHandler: urlHandler,
	}
}
