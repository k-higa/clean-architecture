package stores

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type TransactionManger struct {
	tx *gorm.DB
}

func NewTransactionManger(tx *gorm.DB) *TransactionManger {
	return &TransactionManger{tx: tx}
}

func (t TransactionManger) Begin(ctx context.Context) any {
	return t.tx.Begin()
}

func (t TransactionManger) Commit(ctx context.Context) error {
	return t.tx.Commit().Error
}

func (t TransactionManger) Rollback(ctx context.Context) (any, error) {
	return t.tx.Rollback(), nil
}
func (t TransactionManger) Tx() context.Context {
	return context.WithValue(context.Background(), "tx", t.tx)
}

func (t TransactionManger) Transact(f func(ctx context.Context) error) error {
	tx := t.tx.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	ctx := context.WithValue(context.Background(), "tx", tx)
	err := f(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
