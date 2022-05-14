package repository

import (
	"clean-architecture/usecases"
)

type TransactionManager interface {
	Begin(ctx usecases.Context) interface{}
	Commit(ctx usecases.Context) error
	Rollback(ctx usecases.Context) (interface{}, error)
	Transact(f func(ctx usecases.Context) error) error
	Tx() usecases.Context
}
