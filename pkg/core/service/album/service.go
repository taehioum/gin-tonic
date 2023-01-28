package album

import (
	"context"

	"github.com/taehioum/gin-tonic/pkg/core/dto"
	entity "github.com/taehioum/gin-tonic/pkg/core/entity/album"
	"github.com/taehioum/gin-tonic/pkg/core/port"
)

type Service struct {
	albumRepo port.AlbumRepository
}

var _ port.AlbumService = &Service{}

func (s *Service) GetAlbums(ctx context.Context) []dto.Album {
	albs, _ := s.albumRepo.GetAlbums(ctx)

	res := make([]dto.Album, 0, len(entity.Albums))

	for _, alb := range albs {
		res = append(res, *dto.FromAlbum(alb))
	}

	return res
}

func (*Service) GetAlbum(ctx context.Context, id string) (alb *dto.Album) {
	for _, a := range entity.Albums {
		if a.ID == id {
			return dto.FromAlbum(&a)
		}
	}
	return nil
}

func (*Service) AddAlbum(ctx context.Context, req dto.AlbumCreateRequest) error {

	album := entity.Album{
		ID:     req.ID,
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}

	entity.Albums = append(entity.Albums, album)
	return nil
}
