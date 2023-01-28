package album

import (
	"context"
	"strconv"

	"github.com/taehioum/gin-tonic/pkg/core/dto"
	model "github.com/taehioum/gin-tonic/pkg/core/model/album"
	"github.com/taehioum/gin-tonic/pkg/core/port"
)

type Service struct {
	albumRepo port.AlbumRepository
}

func New(repo port.AlbumRepository) *Service {
	return &Service{albumRepo: repo}
}

func (s *Service) GetAlbums(ctx context.Context) []dto.Album {
	albs, _ := s.albumRepo.GetAlbums(ctx)

	res := make([]dto.Album, 0, len(albs))

	for _, alb := range albs {
		res = append(res, *dto.FromAlbum(alb))
	}

	return res
}

func (s *Service) GetAlbum(ctx context.Context, id string) *dto.Album {
	i, _ := strconv.Atoi(id)
	alb, _ := s.albumRepo.GetAlbumById(ctx, uint(i))
	return dto.FromAlbum(alb)
}

func (s *Service) AddAlbum(ctx context.Context, req dto.AlbumCreateRequest) error {

	album := model.Album{
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}

	s.albumRepo.Save(ctx, &album)
	return nil
}

var _ port.AlbumService = &Service{}
