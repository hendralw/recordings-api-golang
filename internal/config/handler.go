package config

import (
	albumHandler "postgres/internal/handler/album"
	artistHandler "postgres/internal/handler/artist"
	songHandler "postgres/internal/handler/song"
	albumUsecase "postgres/internal/usecase/album"
	artistUsecase "postgres/internal/usecase/artist"
	songUsecase "postgres/internal/usecase/song"
)

type Handler struct {
	AlbumHandler  albumHandler.AlbumHandler
	SongHandler   songHandler.SongHandler
	ArtistHandler artistHandler.ArtistHandler
}

// Function to initialize handler
func InitHandler(albumUsecase albumUsecase.AlbumUsecase, songUsecase songUsecase.SongUsecase, artistUsecase artistUsecase.ArtistUsecase) Handler {
	return Handler{
		AlbumHandler:  albumHandler.NewAlbumHandler(albumUsecase),
		SongHandler:   songHandler.NewSongHandler(songUsecase),
		ArtistHandler: artistHandler.NewArtistHandler(artistUsecase),
	}
}
