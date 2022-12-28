package psql

import (
	"context"
	"database/sql"

	"postgres/internal/entity"

	"github.com/google/uuid"
)

type ArtistPostgres interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (uuid.UUID, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]uuid.UUID, error)
	Update(ctx context.Context, artist entity.Artist) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type artistConnection struct {
	db *sql.DB
}

// The function is to initialize the artist psql repository
func NewArtistPostgres(db *sql.DB) ArtistPostgres {
	return &artistConnection{db: db}
}
