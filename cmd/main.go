package main

import (
	"github.com/gin-gonic/gin"
	"github.com/taehioum/gin-tonic/pkg/adapter/sqlite"
	controller "github.com/taehioum/gin-tonic/pkg/controller/http"
	"github.com/taehioum/gin-tonic/pkg/core/service/album"
)

func main() {
	router := gin.Default()

	db := sqlite.New()

	albumRepo := sqlite.NewAlbumRepository(db)

	albumController := controller.AlbumController{
		AlbumSvc: album.New(albumRepo),
	}

	router.GET("/albums", albumController.GetAlbums)
	router.GET("/albums/:id", albumController.GetAlbum)
	router.POST("/albums", albumController.CreateAlbum)

	router.Run("localhost:8080")
}
