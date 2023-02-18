package sqlite

import (
	"context"

	model "github.com/taehioum/gin-tonic/pkg/core/model/album"
)

type AlbumRepository struct {
	store *Store
}

func NewAlbumRepository(store *Store) *AlbumRepository {
	store.db.AutoMigrate(&model.Album{})
	return &AlbumRepository{store}
}

// GetAlbumById implements port.AlbumRepository
func (r *AlbumRepository) GetAlbumById(ctx context.Context, ID uint) (*model.Album, error) {
	var album model.Album
	result := r.store.
		getTx(ctx).
		Find(&album, ID)

	return &album, result.Error
}

// GetAlbums implements port.AlbumRepository
func (r *AlbumRepository) GetAlbums(ctx context.Context) ([]*model.Album, error) {
	var albums []model.Album
	result := r.store.
		getTx(ctx).
		Find(&albums)

	res := make([]*model.Album, 0)

	for i := range albums {
		res = append(res, &albums[i])
	}
	return res, result.Error
}

// Save implements port.AlbumRepository
func (r *AlbumRepository) Save(ctx context.Context, alb *model.Album) (*model.Album, error) {
	result := r.store.
		getTx(ctx).
		Create(alb)

	return alb, result.Error
}
