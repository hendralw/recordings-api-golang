package psql

import (
	"context"
	"fmt"
	"log"
	"time"

	"postgres/internal/entity"

	"github.com/google/uuid"
)

// Create is function to create artist to database
func (repo *artistConnection) Create(ctx context.Context, artist *entity.Artist) (uuid.UUID, error) {
	// The query insert
	paramID := uuid.New()

	query := `INSERT INTO public."Artists"
		("ArtistID", "Name")
		VALUES($1, $2)
		RETURNING "ArtistID";
		`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, paramID, artist.Name).Scan(&artist.ArtistID)
	if err != nil {
		return uuid.Nil, err
	}

	return artist.ArtistID, nil
}

// Get is function to get specific artist by id from database
func (repo *artistConnection) Get(ctx context.Context, id uuid.UUID) (*entity.Artist, error) {
	// The query select
	query := `
		SELECT "ArtistID", "Name"
		FROM public."Artists"
		WHERE "ArtistID" = $1
		`

	var artist entity.Artist

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query and find the specific artist and then set the result to artist variable
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&artist.ArtistID,
		&artist.Name,
	)
	// If any error
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

// GetAllAlbum is function to get all artists from database
func (repo *artistConnection) GetAllArtist(ctx context.Context) ([]entity.Artist, error) {
	// The query select
	query := `
		SELECT "ArtistID", "Name"
		FROM "Artists"`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var artists []entity.Artist

	// Run the query to get all albums
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the albums is not empty
	for rows.Next() {
		var artist entity.Artist

		// Set to the album variable
		err := rows.Scan(
			&artist.ArtistID,
			&artist.Name,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the album to artists variable
		artists = append(artists, artist)
	}

	return artists, nil
}

// BatchCreate is function to insert some albums in once to database
func (repo *artistConnection) BatchCreate(ctx context.Context, artists []entity.Artist) ([]uuid.UUID, error) {
	var IDs []uuid.UUID

	// Begin transaction
	tx, err := repo.db.Begin()
	if err != nil {
		return IDs, nil
	}
	// If any error, the transaction will be rollback
	defer tx.Rollback()

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query insert
	query := `INSERT INTO public."Artists" 
			("ArtistID", "Name")
			VALUES ($1, $2) 
			RETURNING "ArtistID"`

	// Loop every album
	for _, artist := range artists {
		id := uuid.New()

		// Run query insert of every album to database
		err := tx.QueryRowContext(ctx, query, id, artist.Name).Scan(&id)
		if err != nil {
			log.Printf("error execute insert err: %v", err)
			continue
		}

		// Add the new id to IDs variable
		IDs = append(IDs, id)
	}

	// Commit the transaction
	err = tx.Commit()
	// If any error
	if err != nil {
		return IDs, err
	}

	return IDs, nil
}

// Update is function to update artist in database
func (repo *artistConnection) Update(ctx context.Context, artist entity.Artist) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE "Artists" set "Name"=$1 WHERE "ArtistID"=$2`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, artist.Name, artist.ArtistID)
	if err != nil {
		return err
	}

	// Get how many data has been updated
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected update : %d", rows)
	return nil
}

// Delete is function to delete album in database
func (repo *artistConnection) Delete(ctx context.Context, id uuid.UUID) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from public."Artists" WHERE "ArtistID"=$1`

	// Run the delete query
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Get how many data has been deleted
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected delete : %d", rows)
	return nil
}
