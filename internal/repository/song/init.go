package repository

import (
	"context"
	"database/sql"

	"postgres/internal/entity"
	"postgres/internal/repository/song/cache"
	"postgres/internal/repository/song/psql"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type SongRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Song, error)
	Create(ctx context.Context, song *entity.Song) (uuid.UUID, error)
	GetAllSong(ctx context.Context) ([]entity.Song, error)
	GetAllSongByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Song, error)
	GetAllSongByAlbumID(ctx context.Context, id uuid.UUID) ([]entity.Song, error)
	BatchCreate(ctx context.Context, songs []entity.Song) ([]uuid.UUID, error)
	Update(ctx context.Context, song entity.Song) error
	Delete(ctx context.Context, id uuid.UUID) error

	GetSongCache(ctx context.Context, id uuid.UUID) (*entity.Song, error)
	GetAllSongCache(ctx context.Context) ([]entity.Song, error)
	SetSongCache(ctx context.Context, id uuid.UUID, song entity.Song) error
	SetAllSongCache(ctx context.Context, songs []entity.Song) error
	DeleteSongCache(ctx context.Context, id uuid.UUID) error
}

type songRepository struct {
	postgres psql.SongPostgres
	cache    cache.SongPostgres
}

func NewSongRepository(db *sql.DB, client *redis.Client) SongRepository {
	return &songRepository{
		postgres: psql.NewSongPostgres(db),
		cache:    cache.NewSongRedis(client),
	}
}
