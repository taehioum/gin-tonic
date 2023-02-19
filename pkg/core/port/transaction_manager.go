package port

import (
	"context"
	"database/sql"
)

type TrasactionManager interface {
	WithTx(context.Context, func(context.Context) error) error
	BeginTx(context.Context) (Transaction, error)
	BeginTxWithOptions(context.Context, sql.TxOptions) (Transaction, error)
}

type Transaction interface {
	Ctx() context.Context
	Commit() error
	Rollback() error
}
