package port

import (
	"context"

	entity "github.com/taehioum/gin-tonic/pkg/core/entity/album"
)

type AlbumRepository interface {
	Save(context.Context, *entity.Album) (*entity.Album, error)

	GetAlbumById(context.Context, string) (*entity.Album, error)
	GetAlbums(context.Context) ([]*entity.Album, error)
}
