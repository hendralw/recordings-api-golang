package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"postgres/internal/entity"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Get Specific song cache
func (repo *songConnection) GetSong(ctx context.Context, id uuid.UUID) (*entity.Song, error) {
	var song entity.Song

	key := fmt.Sprintf(songDetailKey, id)

	songsString, err := repo.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return &song, nil
	}
	if err != nil {
		return &song, err
	}

	err = json.Unmarshal([]byte(songsString), &song)
	if err != nil {
		return &song, err
	}

	return &song, nil
}

// GetAllSong is function to get all songs from database
func (repo *songConnection) GetAllSong(ctx context.Context) ([]entity.Song, error) {
	var songs []entity.Song

	songsString, err := repo.client.Get(ctx, songsKey).Result()
	if err == redis.Nil {
		return songs, nil
	}
	if err != nil {
		return songs, err
	}

	err = json.Unmarshal([]byte(songsString), &songs)
	if err != nil {
		return songs, err
	}

	return songs, nil
}

func (repo *songConnection) SetSong(ctx context.Context, id uuid.UUID, song entity.Song) error {
	key := fmt.Sprintf(songDetailKey, id)

	songsString, err := json.Marshal(song)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, key, songsString, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *songConnection) SetAllSong(ctx context.Context, songs []entity.Song) error {
	songsString, err := json.Marshal(songs)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, songsKey, songsString, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *songConnection) Delete(ctx context.Context, id uuid.UUID) error {
	key := fmt.Sprintf(songDetailKey, id)

	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
