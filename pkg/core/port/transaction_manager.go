package port

import (
	"context"
)

type TrasactionManager interface {
	WithTx(context.Context, func(context.Context) error) error
}
