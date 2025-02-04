package storage

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	client *redis.Client
}

func NewStorage() *Storage {
	return &Storage{
		client: redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (s *Storage) Set(ctx context.Context, key string, value []byte) error {
	return s.client.Set(ctx, key, value, 0).Err()
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	return s.client.Get(ctx, key).Result()
}

func (s *Storage) Ping() (string, error) {
	return s.client.Ping(context.Background()).Result()
}
