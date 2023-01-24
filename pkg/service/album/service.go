package album

import (
	"context"

	"github.com/taehioum/gin-tonic/pkg/entity/album"
)

type Service struct {
}

func (*Service) GetAlbums(ctx context.Context) []album.Album {
	return album.Albums
}

func (*Service) GetAlbum(ctx context.Context, id string) (alb *album.Album) {
	for _, a := range album.Albums {
		if a.ID == id {
			return &a
		}
	}
	return nil
}

func (*Service) AddAlbum(ctx context.Context, alb album.Album) error {
	album.Albums = append(album.Albums, alb)
	return nil
}
