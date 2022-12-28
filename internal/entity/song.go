package entity

import "github.com/google/uuid"

// The entity will be used for song definition
type Song struct {
	SongID    uuid.UUID `json:"song_id"`
	Title     string    `json:"title"`
	Lyrics    string    `json:"lyrics"`
	DataAlbum DataAlbum `json:"albums`
}
