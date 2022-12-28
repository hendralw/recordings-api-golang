package cache

import (
	"context"

	"postgres/internal/entity"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type ArtistPostgres interface {
	GetArtist(ctx context.Context, id uuid.UUID) (*entity.Artist, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	SetArtist(ctx context.Context, id uuid.UUID, artist entity.Artist) error
	SetAllArtist(ctx context.Context, artists []entity.Artist) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type artistConnection struct {
	client *redis.Client
}

// The function is to initialize the artist psql repository
func NewArtistRedis(cache *redis.Client) ArtistPostgres {
	return &artistConnection{client: cache}
}
