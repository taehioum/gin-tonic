package port

import (
	"context"

	"github.com/taehioum/gin-tonic/pkg/core/dto"
)

type AlbumService interface {
	GetAlbums(context.Context) []dto.Album
	GetAlbum(context.Context, string) (*dto.Album, error)
	AddAlbum(context.Context, dto.AlbumCreateRequest) error
}
