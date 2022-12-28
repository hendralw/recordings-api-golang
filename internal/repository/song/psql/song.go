package psql

import (
	"context"
	"fmt"
	"log"
	"time"

	"postgres/internal/entity"

	"github.com/google/uuid"
)

// Create is function to create song to database
func (repo *songConnection) Create(ctx context.Context, song *entity.Song) (uuid.UUID, error) {
	// The query insert
	paramID := uuid.New()

	query := `INSERT INTO public."Songs"
		("SongID", "AlbumID", "Title", "Lyrics")
		VALUES($1, $2, $3, $4)
		RETURNING "SongID";
		`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query, paramID, song.DataAlbum.AlbumID, song.Title, song.Lyrics).Scan(&song.SongID)
	if err != nil {
		return uuid.Nil, err
	}

	return song.SongID, nil
}

// Get is function to get specific album by id from database
func (repo *songConnection) Get(ctx context.Context, id uuid.UUID) (*entity.Song, error) {
	// The query select
	fmt.Printf(id.String())
	query := `
			select s."SongID", s."Title", s."Lyrics", s."AlbumID", a."Title", a."ArtistID", a2."Name" 
			from "Songs" s
			inner join "Albums" a on a."AlbumID" = s."AlbumID" 
			inner join "Artists" a2 on a2."ArtistID" = a."ArtistID"
			WHERE "SongID" = $1`

	var song entity.Song

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query and find the specific song and then set the result to album variable
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&song.SongID,
		&song.Title,
		&song.Lyrics,
		&song.DataAlbum.AlbumID,
		&song.DataAlbum.Title,
		&song.DataAlbum.DataArtist.ArtistID,
		&song.DataAlbum.DataArtist.Name,
	)
	// If any error
	if err != nil {
		return nil, err
	}

	return &song, nil
}

// GetAllAlbum is function to get all albums from database
func (repo *songConnection) GetAllSong(ctx context.Context) ([]entity.Song, error) {
	// The query select
	query := `
			select s."SongID", s."Title", s."Lyrics", s."AlbumID", a."Title", a."ArtistID", a2."Name" 
			from "Songs" s
			inner join "Albums" a on a."AlbumID" = s."AlbumID" 
			inner join "Artists" a2 on a2."ArtistID" = a."ArtistID"`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var songs []entity.Song

	// Run the query to get all songs
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the songs is not empty
	for rows.Next() {
		var song entity.Song

		// Set to the song variable
		err := rows.Scan(
			&song.SongID,
			&song.Title,
			&song.Lyrics,
			&song.DataAlbum.AlbumID,
			&song.DataAlbum.Title,
			&song.DataAlbum.DataArtist.ArtistID,
			&song.DataAlbum.DataArtist.Name,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the song to songs variable
		songs = append(songs, song)
	}

	return songs, nil
}

// GetAllAlbum is function to get all albums from database
func (repo *songConnection) GetAllSongByAlbumID(ctx context.Context, id uuid.UUID) ([]entity.Song, error) {
	// The query select
	query := `
			select s."SongID", s."Title", s."Lyrics", s."AlbumID", a."Title", a."ArtistID", a2."Name" 
			from "Songs" s
			inner join "Albums" a on a."AlbumID" = s."AlbumID" 
			inner join "Artists" a2 on a2."ArtistID" = a."ArtistID"
			where a."AlbumID" = $1`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var songs []entity.Song

	// Run the query to get all songs
	rows, err := repo.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the songs is not empty
	for rows.Next() {
		var song entity.Song

		// Set to the song variable
		err := rows.Scan(
			&song.SongID,
			&song.Title,
			&song.Lyrics,
			&song.DataAlbum.AlbumID,
			&song.DataAlbum.Title,
			&song.DataAlbum.DataArtist.ArtistID,
			&song.DataAlbum.DataArtist.Name,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the song to songs variable
		songs = append(songs, song)
	}

	return songs, nil
}

// GetAllAlbum is function to get all albums from database
func (repo *songConnection) GetAllSongByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Song, error) {
	// The query select
	query := `
			select s."SongID", s."Title", s."Lyrics", s."AlbumID", a."Title", a."ArtistID", a2."Name" 
			from "Songs" s
			inner join "Albums" a on a."AlbumID" = s."AlbumID" 
			inner join "Artists" a2 on a2."ArtistID" = a."ArtistID"
			where a."ArtistID" = $1`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var songs []entity.Song

	// Run the query to get all songs
	rows, err := repo.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// If the songs is not empty
	for rows.Next() {
		var song entity.Song

		// Set to the song variable
		err := rows.Scan(
			&song.SongID,
			&song.Title,
			&song.Lyrics,
			&song.DataAlbum.AlbumID,
			&song.DataAlbum.Title,
			&song.DataAlbum.DataArtist.ArtistID,
			&song.DataAlbum.DataArtist.Name,
		)
		// If any error
		if err != nil {
			return nil, err
		}

		// add the song to songs variable
		songs = append(songs, song)
	}

	return songs, nil
}

// BatchCreate is function to insert some albums in once to database
func (repo *songConnection) BatchCreate(ctx context.Context, albums []entity.Song) ([]uuid.UUID, error) {
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
	query := `INSERT INTO public."Songs" 
			("SongID", "AlbumID", "Title", "Lyrics")
			VALUES ($1, $2, $3, $4) 
			RETURNING "SongID"`

	// Loop every album
	for _, album := range albums {
		id := uuid.New()

		// Run query insert of every album to database
		err := tx.QueryRowContext(ctx, query, id, album.DataAlbum.AlbumID, album.Title, album.Lyrics).Scan(&id)
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
func (repo *songConnection) Update(ctx context.Context, song entity.Song) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query update
	query := `UPDATE "Songs" SET "AlbumID"=$1, "Title"=$2, "Lyrics"=$3 WHERE "SongID"=$4`

	// Run the query
	result, err := repo.db.ExecContext(ctx, query, song.DataAlbum.AlbumID, song.Title, song.Lyrics, song.SongID)
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
func (repo *songConnection) Delete(ctx context.Context, id uuid.UUID) error {
	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// The query delete
	query := `DELETE from public."Songs" WHERE "SongID"=$1`

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
