package usecase

import (
	"context"

	"postgres/internal/entity"
	songRepository "postgres/internal/repository/song"

	"github.com/google/uuid"
)

type SongUsecase interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Song, error)
	Create(ctx context.Context, song *entity.Song) (*entity.Song, error)
	GetAllSong(ctx context.Context) ([]entity.Song, error)
	GetAllSongByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Song, error)
	GetAllSongByAlbumID(ctx context.Context, id uuid.UUID) ([]entity.Song, error)
	BatchCreate(ctx context.Context, songs []entity.Song) ([]entity.Song, error)
	Update(ctx context.Context, song entity.Song) (entity.Song, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type songUsecase struct {
	songRepository songRepository.SongRepository
}

// The function is to initialize the song usecase
func NewSongUsecase(songRepository songRepository.SongRepository) SongUsecase {
	return &songUsecase{
		songRepository: songRepository,
	}
}
