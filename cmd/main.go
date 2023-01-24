package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/taehioum/gin-tonic/pkg/controller/http"
	"github.com/taehioum/gin-tonic/pkg/service/album"
)

func main() {
	router := gin.Default()

	albumController := controller.AlbumController{
		AlbumSvc: album.Service{},
	}

	router.GET("/albums", albumController.GetAlbums)
	router.GET("/albums/:id", albumController.GetAlbum)
	router.POST("/albums", albumController.CreateAlbum)

	router.Run("localhost:8080")
}
