package shortener

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/EgorSalenko/tiny/storage"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Service struct {
	storage *storage.Storage
}

func NewService(s *storage.Storage) *Service {
	return &Service{storage: s}
}

func (r *Service) Hash(ctx context.Context, url string) (*storage.Tiny, error) {
	hash := md5Hash(url)

	// Attempt to fetch from cache
	result, err := r.storage.Get(ctx, hash)
	if err == nil { // Cache hit
		var data storage.Tiny
		if err := json.Unmarshal([]byte(result), &data); err != nil {
			log.Error().Err(err).Str("url", url).Msg("Failed to unmarshal data")
			return nil, err
		}
		return &data, nil
	}

	if err != redis.Nil { // Unexpected error
		log.Error().Err(err).Str("url", url).Msg("Failed to fetch from Redis")
		return nil, err
	}

	// Cache miss, create new Tiny entry
	data := storage.Tiny{URL: url, Hash: hash}

	raw, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("Failed to marshal data")
		return nil, err
	}

	// Store new data in Redis
	if err := r.storage.Set(ctx, hash, raw); err != nil {
		log.Error().Err(err).Str("url", url).Msg("Failed to save hash")
		return nil, err
	}

	log.Info().Str("url", url).Msg("Hash successfully saved")
	return &data, nil

}

func (r *Service) GetUrlByHash(ctx context.Context, hash string) (url string, err error) {
	raw, err := r.storage.Get(ctx, hash)
	var data storage.Tiny
	err = json.Unmarshal([]byte(raw), &data)
	if err != nil {
		return
	}
	return data.URL, nil
}

func md5Hash(src string) string {
	hasher := md5.New()
	hasher.Write([]byte(src))
	return hex.EncodeToString(hasher.Sum(nil))
}
