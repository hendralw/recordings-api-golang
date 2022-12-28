package cache

import (
	"context"

	"postgres/internal/entity"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type SongPostgres interface {
	GetSong(ctx context.Context, id uuid.UUID) (*entity.Song, error)
	GetAllSong(ctx context.Context) ([]entity.Song, error)
	SetSong(ctx context.Context, id uuid.UUID, song entity.Song) error
	SetAllSong(ctx context.Context, songs []entity.Song) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type songConnection struct {
	client *redis.Client
}

// The function is to initialize the song psql repository
func NewSongRedis(cache *redis.Client) SongPostgres {
	return &songConnection{client: cache}
}
