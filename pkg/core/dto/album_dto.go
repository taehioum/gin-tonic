package dto

import entity "github.com/taehioum/gin-tonic/pkg/core/entity/album"

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func FromAlbum(alb *entity.Album) *Album {

	return &Album{
		ID:     alb.ID,
		Title:  alb.Title,
		Artist: alb.Artist,
		Price:  alb.Price,
	}
}
