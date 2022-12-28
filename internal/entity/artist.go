package entity

import "github.com/google/uuid"

// The entity will be used for album definition
type Artist struct {
	ArtistID uuid.UUID `json:"artist_id"`
	Name     string    `json:"name"`
}
