package sqlite

import (
	"context"

	model "github.com/taehioum/gin-tonic/pkg/core/model/album"
	"github.com/taehioum/gin-tonic/pkg/core/port"
	"gorm.io/gorm"
)

type AlbumRepository struct {
	db *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *AlbumRepository {
	db.AutoMigrate(&model.Album{})
	return &AlbumRepository{db: db}
}

// GetAlbumById implements port.AlbumRepository
func (r *AlbumRepository) GetAlbumById(ctx context.Context, ID uint) (*model.Album, error) {
	var album model.Album
	result := r.db.Find(&album, ID)

	return &album, result.Error
}

// GetAlbums implements port.AlbumRepository
func (r *AlbumRepository) GetAlbums(ctx context.Context) ([]*model.Album, error) {
	var albums []model.Album
	result := r.db.Find(&albums)

	res := make([]*model.Album, 0)

	for i := range albums {
		res = append(res, &albums[i])
	}
	return res, result.Error
}

// Save implements port.AlbumRepository
func (r *AlbumRepository) Save(ctx context.Context, alb *model.Album) (*model.Album, error) {
	result := r.db.Create(alb)

	return alb, result.Error
}

var _ port.AlbumRepository = &AlbumRepository{}
