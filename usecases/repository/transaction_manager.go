package repository

import "golang.org/x/net/context"

type TransactionManager interface {
	Begin(ctx context.Context) interface{}
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) (interface{}, error)
	Transact(f func(ctx context.Context) error) error
	Tx() context.Context
}
