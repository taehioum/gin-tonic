package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taehioum/gin-tonic/pkg/entity/album"
	albumSvc "github.com/taehioum/gin-tonic/pkg/service/album"
)

type AlbumController struct {
	AlbumSvc albumSvc.Service
}

func (ctl *AlbumController) GetAlbums(c *gin.Context) {

	ctx := c.Request.Context()
	res := ctl.AlbumSvc.GetAlbums(ctx)
	c.JSON(http.StatusOK, res)
}

func (ctl *AlbumController) CreateAlbum(c *gin.Context) {

	var newAlbum album.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	_ = ctl.AlbumSvc.AddAlbum(c.Request.Context(), newAlbum)
	c.Status(http.StatusCreated)
}

func (ctl *AlbumController) GetAlbum(c *gin.Context) {

	id := c.Param("id")

	alb := ctl.AlbumSvc.GetAlbum(c.Request.Context(), id)
	if alb == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, alb)
}
