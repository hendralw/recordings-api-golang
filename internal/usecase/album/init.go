package usecase

import (
	"context"

	"postgres/internal/entity"
	albumRepository "postgres/internal/repository/album"

	"github.com/google/uuid"
)

type AlbumUsecase interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Album, error)
	Create(ctx context.Context, album *entity.Album) (*entity.Album, error)
	GetAllAlbum(ctx context.Context) ([]entity.Album, error)
	GetAllAlbumByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Album, error)
	BatchCreate(ctx context.Context, albums []entity.Album) ([]entity.Album, error)
	Update(ctx context.Context, album entity.Album) (entity.Album, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type albumUsecase struct {
	albumRepository albumRepository.AlbumRepository
}

// The function is to initialize the album usecase
func NewAlbumUsecase(albumRepository albumRepository.AlbumRepository) AlbumUsecase {
	return &albumUsecase{
		albumRepository: albumRepository,
	}
}
