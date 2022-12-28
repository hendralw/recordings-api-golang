package repository

import (
	"context"

	"postgres/internal/entity"

	"github.com/google/uuid"
)

// It will call the function Create in psql/album
func (repo *albumRepository) Create(ctx context.Context, album *entity.Album) (uuid.UUID, error) {
	return repo.postgres.Create(ctx, album)
}

// It will call the function Get in psql/album
func (repo *albumRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Album, error) {
	return repo.postgres.Get(ctx, id)
}

// It will call the function GetAllAlbum in psql/album
func (repo *albumRepository) GetAllAlbum(ctx context.Context) ([]entity.Album, error) {
	return repo.postgres.GetAllAlbum(ctx)
}

func (repo *albumRepository) GetAllAlbumByArtistID(ctx context.Context, id uuid.UUID) ([]entity.Album, error) {
	return repo.postgres.GetAllAlbumByArtistID(ctx, id)
}

// It will call the function BatchCreate in psql/album
func (repo *albumRepository) BatchCreate(ctx context.Context, albums []entity.Album) ([]uuid.UUID, error) {
	return repo.postgres.BatchCreate(ctx, albums)
}

// It will call the function Update in psql/album
func (repo *albumRepository) Update(ctx context.Context, album entity.Album) error {
	return repo.postgres.Update(ctx, album)
}

// It will call the function Delete in psql/album
func (repo *albumRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return repo.postgres.Delete(ctx, id)
}

func (repo *albumRepository) GetAlbumCache(ctx context.Context, id uuid.UUID) (*entity.Album, error) {
	return repo.cache.GetAlbum(ctx, id)
}

func (repo *albumRepository) GetAllAlbumCache(ctx context.Context) ([]entity.Album, error) {
	return repo.cache.GetAllAlbum(ctx)
}

func (repo *albumRepository) SetAlbumCache(ctx context.Context, id uuid.UUID, album entity.Album) error {
	return repo.cache.SetAlbum(ctx, id, album)
}

func (repo *albumRepository) SetAllAlbumCache(ctx context.Context, albums []entity.Album) error {
	return repo.cache.SetAllAlbum(ctx, albums)
}

func (repo *albumRepository) DeleteAlbumCache(ctx context.Context, id uuid.UUID) error {
	return repo.cache.Delete(ctx, id)
}
