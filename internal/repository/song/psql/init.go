package psql

import (
	"context"
	"database/sql"

	"postgres/internal/entity"

	"github.com/google/uuid"
)

type SongPostgres interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Song, error)
	Create(ctx context.Context, song *entity.Song) (uuid.UUID, error)
	GetAllSong(ctx context.Context) ([]entity.Song, error)
	GetAllSongByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Song, error)
	GetAllSongByAlbumID(ctx context.Context, id uuid.UUID) ([]entity.Song, error)
	BatchCreate(ctx context.Context, songs []entity.Song) ([]uuid.UUID, error)
	Update(ctx context.Context, song entity.Song) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type songConnection struct {
	db *sql.DB
}

// The function is to initialize the song psql repository
func NewSongPostgres(db *sql.DB) SongPostgres {
	return &songConnection{db: db}
}
