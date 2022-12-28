package psql

import (
	"context"
	"database/sql"

	"postgres/internal/entity"

	"github.com/google/uuid"
)

type AlbumPostgres interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (uuid.UUID, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	GetAllAlbumByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]uuid.UUID, error)
	Update(ctx context.Context, album entity.Album) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type albumConnection struct {
	db *sql.DB
}

// The function is to initialize the album psql repository
func NewAlbumPostgres(db *sql.DB) AlbumPostgres {
	return &albumConnection{db: db}
}
