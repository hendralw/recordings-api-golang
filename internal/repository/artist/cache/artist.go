package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"postgres/internal/entity"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

// Get Specific artist cache
func (repo *artistConnection) GetArtist(ctx context.Context, id uuid.UUID) (*entity.Artist, error) {
	var artist entity.Artist

	key := fmt.Sprintf(artistDetailKey, id)

	artistString, err := repo.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return &artist, nil
	}
	if err != nil {
		return &artist, err
	}

	err = json.Unmarshal([]byte(artistString), &artist)
	if err != nil {
		return &artist, err
	}

	return &artist, nil
}

// GetAllArtist is function to get all artists from database
func (repo *artistConnection) GetAllArtist(ctx context.Context) ([]entity.Artist, error) {
	var artists []entity.Artist

	artistString, err := repo.client.Get(ctx, artistsKey).Result()
	if err == redis.Nil {
		return artists, nil
	}
	if err != nil {
		return artists, err
	}

	err = json.Unmarshal([]byte(artistString), &artists)
	if err != nil {
		return artists, err
	}

	return artists, nil
}

func (repo *artistConnection) SetArtist(ctx context.Context, id uuid.UUID, artist entity.Artist) error {
	key := fmt.Sprintf(artistDetailKey, id)

	artistString, err := json.Marshal(artist)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, key, artistString, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *artistConnection) SetAllArtist(ctx context.Context, artists []entity.Artist) error {
	artistString, err := json.Marshal(artists)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, artistsKey, artistString, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *artistConnection) Delete(ctx context.Context, id uuid.UUID) error {
	key := fmt.Sprintf(artistDetailKey, id)

	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
