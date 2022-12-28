package usecase

import (
	"context"
	"fmt"

	"postgres/internal/entity"

	"github.com/google/uuid"
)

// // It will call the function Get in song repository
func (usecase *songUsecase) Get(ctx context.Context, id uuid.UUID) (*entity.Song, error) {
	// Get from cache
	song, err := usecase.songRepository.GetSongCache(ctx, id)
	if err != nil {
		return song, err
	}

	if song.SongID != uuid.Nil {
		return song, nil
	}

	// Get from db
	song, err = usecase.songRepository.Get(ctx, id)
	if err != nil {
		return song, err
	}

	// Set to cache
	if err = usecase.songRepository.SetSongCache(ctx, id, *song); err != nil {
		return song, err
	}

	return song, nil
}

// It will call the function Create in song repository
func (usecase *songUsecase) Create(ctx context.Context, song *entity.Song) (*entity.Song, error) {
	var newSong *entity.Song

	// Create song
	id, err := usecase.songRepository.Create(ctx, song)
	if err != nil {
		return newSong, err
	}
	fmt.Println(id)

	// Find new song
	newSong, err = usecase.songRepository.Get(ctx, id)
	if err != nil {
		return newSong, err
	}

	// Find all songs
	songs, err := usecase.songRepository.GetAllSong(ctx)
	if err != nil {
		return newSong, err
	}

	// Set to specific cache
	if err = usecase.songRepository.SetSongCache(ctx, id, *newSong); err != nil {
		return newSong, err
	}

	// Set all cache
	if err = usecase.songRepository.SetAllSongCache(ctx, songs); err != nil {
		return newSong, err
	}

	return newSong, nil
}

// // It will call the function Get in song repository
func (usecase *songUsecase) GetAllSongByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Song, error) {
	// Get from db
	song, err := usecase.songRepository.GetAllSongByArtistID(ctx, id)
	if err != nil {
		return song, err
	}

	return song, nil
}

// // It will call the function Get in song repository
func (usecase *songUsecase) GetAllSongByAlbumID(ctx context.Context, id uuid.UUID) ([]entity.Song, error) {
	// Get from db
	song, err := usecase.songRepository.GetAllSongByAlbumID(ctx, id)
	if err != nil {
		return song, err
	}

	return song, nil
}

// It will call the function GetAllSong in song repository
func (usecase *songUsecase) GetAllSong(ctx context.Context) ([]entity.Song, error) {
	var songs []entity.Song

	// // Get from cache
	songs, err := usecase.songRepository.GetAllSongCache(ctx)
	if err != nil {
		return songs, err
	}

	if len(songs) > 0 {
		return songs, nil
	}

	// Get from db
	songs, err = usecase.songRepository.GetAllSong(ctx)
	if err != nil {
		return songs, err
	}

	// Set to cache
	if err = usecase.songRepository.SetAllSongCache(ctx, songs); err != nil {
		return songs, err
	}

	return songs, nil
}

// It will call the function BatchCreate in song repository
func (usecase *songUsecase) BatchCreate(ctx context.Context, songs []entity.Song) ([]entity.Song, error) {
	var newSongs []entity.Song

	// Batch create and get the new id
	ids, err := usecase.songRepository.BatchCreate(ctx, songs)
	if err != nil {
		return newSongs, err
	}

	// Get detail song by ids
	for _, id := range ids {
		// Get from db
		song, err := usecase.songRepository.Get(ctx, id)
		if err != nil {
			return newSongs, err
		}

		// Set to specific cache
		if err = usecase.songRepository.SetSongCache(ctx, id, *song); err != nil {
			return newSongs, err
		}

		newSongs = append(newSongs, *song)
	}

	// Find all albums
	allSongs, err := usecase.songRepository.GetAllSong(ctx)
	if err != nil {
		return newSongs, err
	}

	// Set all cache
	if err = usecase.songRepository.SetAllSongCache(ctx, allSongs); err != nil {
		return newSongs, err
	}

	return newSongs, nil
}

// It will call the function Update in song repository
func (usecase *songUsecase) Update(ctx context.Context, song entity.Song) (entity.Song, error) {
	var updatedSong *entity.Song

	// Update song
	err := usecase.songRepository.Update(ctx, song)
	if err != nil {
		return *updatedSong, err
	}

	// Find new song
	updatedSong, err = usecase.songRepository.Get(ctx, song.SongID)
	if err != nil {
		return *updatedSong, err
	}

	// Find all songs
	songs, err := usecase.songRepository.GetAllSong(ctx)
	if err != nil {
		return *updatedSong, err
	}

	// Set to specific cache
	if err = usecase.songRepository.SetSongCache(ctx, song.SongID, *updatedSong); err != nil {
		return *updatedSong, err
	}

	// Set all cache
	if err = usecase.songRepository.SetAllSongCache(ctx, songs); err != nil {
		return *updatedSong, err
	}

	return *updatedSong, nil
}

// It will call the function Delete in album repository
func (usecase *songUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	// Delete from db
	if err := usecase.songRepository.Delete(ctx, id); err != nil {
		return err
	}

	// Delete from cache
	if err := usecase.songRepository.DeleteSongCache(ctx, id); err != nil {
		return err
	}

	// Find all songs
	songs, err := usecase.songRepository.GetAllSong(ctx)
	if err != nil {
		return err
	}

	// Set all cache
	if err = usecase.songRepository.SetAllSongCache(ctx, songs); err != nil {
		return err
	}

	return nil
}
