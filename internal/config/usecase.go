package config

import (
	albumRepository "postgres/internal/repository/album"
	artistRepository "postgres/internal/repository/artist"
	songRepository "postgres/internal/repository/song"
	albumUsecase "postgres/internal/usecase/album"
	artistUsecase "postgres/internal/usecase/artist"
	songUsecase "postgres/internal/usecase/song"
)

type Usecase struct {
	AlbumUsecase  albumUsecase.AlbumUsecase
	SongUsecase   songUsecase.SongUsecase
	ArtistUsecase artistUsecase.ArtistUsecase
}

// Function to initialize usecase
func InitUsecase(albumRepository albumRepository.AlbumRepository, songRepository songRepository.SongRepository, artistRepository artistRepository.ArtistRepository) Usecase {
	return Usecase{
		AlbumUsecase:  albumUsecase.NewAlbumUsecase(albumRepository),
		SongUsecase:   songUsecase.NewSongUsecase(songRepository),
		ArtistUsecase: artistUsecase.NewArtistUsecase(artistRepository),
	}
}
