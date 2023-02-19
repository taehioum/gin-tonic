package port

import (
	"context"

	model "github.com/taehioum/gin-tonic/pkg/core/model/album"
)

type AlbumRepository interface {
	Save(context.Context, *model.Album) (*model.Album, error)

	GetAlbumByID(context.Context, uint) (*model.Album, error)
	GetAlbums(context.Context) ([]*model.Album, error)
}
