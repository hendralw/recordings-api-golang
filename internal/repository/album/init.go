package repository

import (
	"context"
	"database/sql"

	"postgres/internal/entity"
	"postgres/internal/repository/album/cache"
	"postgres/internal/repository/album/psql"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type AlbumRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (uuid.UUID, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	GetAllAlbumByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]uuid.UUID, error)
	Update(ctx context.Context, album entity.Album) error
	Delete(ctx context.Context, id uuid.UUID) error

	GetAlbumCache(ctx context.Context, id uuid.UUID) (*entity.Album, error)
	GetAllAlbumCache(ctx context.Context) ([]entity.Album, error)
	SetAlbumCache(ctx context.Context, id uuid.UUID, album entity.Album) error
	SetAllAlbumCache(ctx context.Context, albums []entity.Album) error
	DeleteAlbumCache(ctx context.Context, id uuid.UUID) error
}

type albumRepository struct {
	postgres psql.AlbumPostgres
	cache    cache.AlbumPostgres
}

// The function is to initialize the album repository
func NewAlbumRepository(db *sql.DB, client *redis.Client) AlbumRepository {
	return &albumRepository{
		postgres: psql.NewAlbumPostgres(db),
		cache:    cache.NewAlbumRedis(client),
	}
}

// func NewAlbumRepository2(db *sql.DB) AlbumRepository {
// 	return &albumRepository{
// 		postgres: psql.NewAlbumPostgres(db),
// 	}
// }
