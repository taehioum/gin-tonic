package album

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taehioum/gin-tonic/pkg/adapter/sqlite"
	"github.com/taehioum/gin-tonic/pkg/core/dto"
)

var store = sqlite.New()
var albumRepo = sqlite.NewAlbumRepository(store)
var albumSvc = New(albumRepo, store)

func TestCreateAlbum(t *testing.T) {

	t.Run("create new album", func(t *testing.T) {
		req := dto.AlbumCreateRequest{
			Title:  "new album",
			Artist: "jojo",
			Price:  50,
		}
		err := albumSvc.AddAlbum(context.TODO(), req)
		assert.Nil(t, err)
	})

}

func TestGetAlbum(t *testing.T) {

	store := sqlite.New()
	albumRepo := sqlite.NewAlbumRepository(store)
	albumSvc := New(albumRepo, store)

	t.Run("get album with existing id", func(t *testing.T) {
		album, err := albumSvc.GetAlbum(context.TODO(), "1")
		assert.Nil(t, err)
		assert.Equal(t, uint(1), album.ID)
	})

}
