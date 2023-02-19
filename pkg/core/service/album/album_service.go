package album

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/taehioum/gin-tonic/pkg/core/dto"
	model "github.com/taehioum/gin-tonic/pkg/core/model/album"
	"github.com/taehioum/gin-tonic/pkg/core/port"
	"github.com/taehioum/gin-tonic/pkg/pkgerr"
)

type Service struct {
	albumRepo port.AlbumRepository
	tm        port.TrasactionManager
}

func New(repo port.AlbumRepository, tm port.TrasactionManager) *Service {
	return &Service{albumRepo: repo, tm: tm}
}

func (s *Service) GetAlbums(ctx context.Context) []dto.Album {
	albs, _ := s.albumRepo.GetAlbums(ctx)

	res := make([]dto.Album, 0, len(albs))

	for _, alb := range albs {
		res = append(res, *dto.FromAlbum(alb))
	}

	return res
}

func (s *Service) GetAlbum(ctx context.Context, id string) (*dto.Album, error) {
	i, _ := strconv.Atoi(id)
	alb, err := s.albumRepo.GetAlbumByID(ctx, uint(i))
	if errors.Is(err, pkgerr.ErrAlbumNotFound) {
		fmt.Println("not found")
		return &dto.Album{}, err
	}

	return dto.FromAlbum(alb), nil
}

func (s *Service) AddAlbum(ctx context.Context, req dto.AlbumCreateRequest) error {

	album := model.Album{
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}

	err := s.tm.WithTx(ctx, func(context.Context) error {
		_, err := s.albumRepo.Save(ctx, &album)
		return errors.Wrap(err, "failed to add album")
	})
	return err
}

func (s *Service) AddAlbum2(ctx context.Context, req dto.AlbumCreateRequest) error {

	album := model.Album{
		Title:  req.Title,
		Artist: req.Artist,
		Price:  req.Price,
	}

	tx, _ := s.tm.BeginTx(ctx)
	defer tx.Rollback()

	_, err := s.albumRepo.Save(tx.Ctx(), &album)
	if err != nil {
		tx.Commit()
	}
	return err
}
