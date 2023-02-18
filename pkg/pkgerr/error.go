package pkgerr

import "errors"

// 4XX

var (
	ErrAlbumNotFound = errors.New("resource was not found")
)
