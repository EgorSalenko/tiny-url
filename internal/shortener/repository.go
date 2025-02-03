package shortener

import (
	"context"
	"github.com/EgorSalenko/tiny/internal/storage"
)

type Repository struct {
	storage *storage.Storage
}

func NewRepository(s *storage.Storage) Repository {
	return Repository{storage: s}
}

func (r Repository) Hash(ctx context.Context, url string) (string, error) {
	return r.storage.UrlHash(ctx, url)
}
