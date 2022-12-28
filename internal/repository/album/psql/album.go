package psql

import (
	"context"
	"fmt"
	"log"
	"time"

	"postgres/internal/entity"

	"github.com/google/uuid"
)

// Create is function to create album to database
func (repo *albumConnection) Create(ctx context.Context, album *entity.Album) (uuid.UUID, error) {
	// The query insert
	paramID := uuid.New()

	query := `INSERT INTO public."Albums"
		("AlbumID", "ArtistID", "Title", "Price")
		VALUES($1, $2, $3, $4)
		RETURNING "AlbumID";
		`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, paramID, album.DataArtist.ArtistID, album.Title, album.Price).Scan(&album.AlbumID)
	if err != nil {
		return uuid.Nil, err
	}

	return album.AlbumID, nil
}

// Get is function to get specific album by id from database
func (repo *albumConnection) Get(ctx context.Context, id uuid.UUID) (*entity.Album, error) {
	// The query select
	query := `
			select a."AlbumID", a."Title", a."Price", a."ArtistID", a2."Name" from "Albums" a
			inner join "Artists" a2 on a2."ArtistID" = a."ArtistID" 
			where a."AlbumID" = $1`

	var album entity.Album

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query and find the specific album and then set the result to album variable
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&album.AlbumID,
		&album.Title,
		&album.Price,
		&album.DataArtist.ArtistID,
		&album.DataArtist.Name,
	)
	// If any error
	if err != nil {
		return nil, err
	}

	return &album, nil
}

// GetAllAlbum is function to get all albums from database
func (repo *albumConnection) GetAllAlbum(ctx context.Context) ([]entity.Album, error) {
	// The query select
	query := `
			select a."AlbumID", a."Title", a."Price", a."ArtistID", a2."Name" from "Albums" a 
			inner join "Artists" a2 on a2."ArtistID" = a."ArtistID"`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var albums []entity.Album

	// Run the query to get all albums
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the albums is not empty
	for rows.Next() {
		var album entity.Album

		// Set to the album variable
		err := rows.Scan(
			&album.AlbumID,
			&album.Title,
			&album.Price,
			&album.DataArtist.ArtistID,
			&album.DataArtist.Name,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the album to albums variable
		albums = append(albums, album)
	}

	return albums, nil
}

// GetAllAlbum is function to get all albums from database
func (repo *albumConnection) GetAllAlbumByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Album, error) {
	// The query select
	query := `
			select a."AlbumID", a."Title", a."Price", a."ArtistID", a2."Name" from "Albums" a 
			inner join "Artists" a2 on a2."ArtistID" = a."ArtistID" 
			where a."ArtistID" = $1
			`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var albums []entity.Album

	// Run the query to get all albums
	rows, err := repo.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the albums is not empty
	for rows.Next() {
		var album entity.Album

		// Set to the album variable
		err := rows.Scan(
			&album.AlbumID,
			&album.Title,
			&album.Price,
			&album.DataArtist.ArtistID,
			&album.DataArtist.Name,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the album to albums variable
		albums = append(albums, album)
	}

	return albums, nil
}

// BatchCreate is function to insert some albums in once to database
func (repo *albumConnection) BatchCreate(ctx context.Context, albums []entity.Album) ([]uuid.UUID, error) {
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
	query := `INSERT INTO public."Albums" 
			("AlbumID", "ArtistID", "Title", "Price")
			VALUES ($1, $2, $3, $4) 
			RETURNING "AlbumID"`

	// Loop every album
	for _, album := range albums {
		id := uuid.New()

		// Run query insert of every album to database
		err := tx.QueryRowContext(ctx, query, id, album.DataArtist.ArtistID, album.Title, album.Price).Scan(&id)
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

// Update is function to update album in database
func (repo *albumConnection) Update(ctx context.Context, album entity.Album) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE "Albums" set "Title"=$1, "ArtistID"=$2, "Price"=$3 WHERE "AlbumID"=$4`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, album.Title, album.DataArtist.ArtistID, album.Price, album.AlbumID)
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
func (repo *albumConnection) Delete(ctx context.Context, id uuid.UUID) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from public."Albums" WHERE "AlbumID"=$1`

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
