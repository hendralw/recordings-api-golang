package entity

import "github.com/google/uuid"

// The entity will be used for album definition
type Album struct {
	AlbumID    uuid.UUID `json:"album_id"`
	Title      string    `json:"title"`
	Price      float32   `json:"price"`
	DataArtist Artist    `json:"artist"`
}

type DataAlbum struct {
	AlbumID    uuid.UUID `json:"album_id"`
	Title      string    `json:"title"`
	DataArtist Artist    `json:"artist"`
}
