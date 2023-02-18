package sqlite

import (
	"context"
	"database/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

type ctxKeyTransaction struct{}

func New() *Store {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil
	}
	return &Store{db: db}
}

func (s *Store) WithTx(ctx context.Context, fn func(context.Context) error) error {
	tx := s.db.Begin(&sql.TxOptions{})
	defer tx.Rollback()

	txCtx := context.WithValue(ctx, ctxKeyTransaction{}, tx)

	err := fn(txCtx)
	if err == nil {
		tx.Commit()
	}
	return err
}

// if there is an ongoing transaction in current context, return that tx
// return a new tx other wise
func (s *Store) getTx(ctx context.Context) *gorm.DB {
	txAny := ctx.Value(ctxKeyTransaction{})
	if tx, ok := txAny.(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}
	return s.db.WithContext(ctx)
}
