package config

import (
	"database/sql"

	albumRepository "postgres/internal/repository/album"
	artistRepository "postgres/internal/repository/artist"
	songRepository "postgres/internal/repository/song"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	AlbumRepository  albumRepository.AlbumRepository
	SongRepository   songRepository.SongRepository
	ArtistRepository artistRepository.ArtistRepository
}

// Function to initialize repository
func InitRepository(db *sql.DB, cache *redis.Client) Repository {
	return Repository{
		AlbumRepository:  albumRepository.NewAlbumRepository(db, cache),
		SongRepository:   songRepository.NewSongRepository(db, cache),
		ArtistRepository: artistRepository.NewArtistRepository(db, cache),
	}
}
