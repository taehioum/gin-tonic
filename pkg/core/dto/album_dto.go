package dto

import model "github.com/taehioum/gin-tonic/pkg/core/model/album"

type Album struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func FromAlbum(alb *model.Album) *Album {

	return &Album{
		ID:     alb.ID,
		Title:  alb.Title,
		Artist: alb.Artist,
		Price:  alb.Price,
	}
}
