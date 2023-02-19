package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/taehioum/gin-tonic/pkg/core/port"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

type ctxKeyTx struct{}

func New() *Store {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil
	}
	return &Store{db: db}
}

func (s *Store) BeginTx(ctx context.Context) (port.Transaction, error) {
	return s.BeginTxWithOptions(ctx, sql.TxOptions{})
}

func (s *Store) BeginTxWithOptions(ctx context.Context, options sql.TxOptions) (port.Transaction, error) {
	tx := s.db.Begin(&options)
	txCtx := context.WithValue(ctx, ctxKeyTx{}, tx)

	return &transaction{
		ctx:      txCtx,
		commit:   func() error { return tx.Commit().Error },
		rollback: func() error { return tx.Rollback().Error },
	}, nil
}

func (s *Store) WithTx(ctx context.Context, fn func(context.Context) error) error {
	tx := s.db.Begin(&sql.TxOptions{})
	defer tx.Rollback()

	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	txCtx := context.WithValue(ctx, ctxKeyTx{}, tx)

	if err := fn(txCtx); err != nil {
		if rerr := tx.Rollback().Error; rerr != nil {
			err = fmt.Errorf("rollbacking transaction: %v: %w", rerr, err)
		}
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

// if there is an ongoing transaction in current context, return that tx
// return a new tx other wise
func (s *Store) getTx(ctx context.Context) *gorm.DB {
	txAny := ctx.Value(ctxKeyTx{})
	if tx, ok := txAny.(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}
	return s.db.WithContext(ctx)
}
