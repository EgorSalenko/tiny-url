package shortener

import (
	"context"
)

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) GetUrlHash(ctx context.Context, srcUlr string) (string, error) {
	return s.repository.Hash(ctx, srcUlr)
}
