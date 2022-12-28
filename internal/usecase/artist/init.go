package usecase

import (
	"context"

	"postgres/internal/entity"
	artistRepository "postgres/internal/repository/artist"

	"github.com/google/uuid"
)

type ArtistUsecase interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (*entity.Artist, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]entity.Artist, error)
	Update(ctx context.Context, artist entity.Artist) (entity.Artist, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type artistUsecase struct {
	artistRepository artistRepository.ArtistRepository
}

// The function is to initialize the artist usecase
func NewArtistUsecase(artistRepository artistRepository.ArtistRepository) ArtistUsecase {
	return &artistUsecase{
		artistRepository: artistRepository,
	}
}
