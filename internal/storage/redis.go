package storage

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/redis/go-redis/v9"
	"time"
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

func (s *Storage) UrlHash(ctx context.Context, url string) (string, error) {
	r, err := s.client.Get(ctx, url).Result()
	if err == nil {
		return r, nil
	}

	r, err = s.client.Set(ctx, url, md5Hash(url), time.Minute*60).Result()
	if err != nil {
		return "", err
	}
	return r, nil
}

func md5Hash(src string) string {
	hasher := md5.New()
	hasher.Write([]byte(src))
	return hex.EncodeToString(hasher.Sum(nil))
}
