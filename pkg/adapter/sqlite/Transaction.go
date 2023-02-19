package sqlite

import "context"

type transaction struct {
	ctx      context.Context
	commit   func() error
	rollback func() error
}

func (tx *transaction) Ctx() context.Context {
	return tx.ctx
}

func (tx *transaction) Commit() error {
	return tx.commit()
}

func (tx *transaction) Rollback() error {
	return tx.rollback()
}
