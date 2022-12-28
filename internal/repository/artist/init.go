package repository

import (
	"context"
	"database/sql"

	"postgres/internal/entity"
	"postgres/internal/repository/artist/cache"
	"postgres/internal/repository/artist/psql"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type ArtistRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (uuid.UUID, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]uuid.UUID, error)
	Update(ctx context.Context, artist entity.Artist) error
	Delete(ctx context.Context, id uuid.UUID) error

	GetArtistCache(ctx context.Context, id uuid.UUID) (*entity.Artist, error)
	GetAllArtistCache(ctx context.Context) ([]entity.Artist, error)
	SetArtistCache(ctx context.Context, id uuid.UUID, artist entity.Artist) error
	SetAllArtistCache(ctx context.Context, artists []entity.Artist) error
	DeleteArtistCache(ctx context.Context, id uuid.UUID) error
}

type artistRepository struct {
	postgres psql.ArtistPostgres
	cache    cache.ArtistPostgres
}

// The function is to initialize the artist repository
func NewArtistRepository(db *sql.DB, client *redis.Client) ArtistRepository {
	return &artistRepository{
		postgres: psql.NewArtistPostgres(db),
		cache:    cache.NewArtistRedis(client),
	}
}
