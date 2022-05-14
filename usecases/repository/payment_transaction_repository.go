package repository

import (
	"clean-architecture/domains"
	"golang.org/x/net/context"
)

type PaymentTransactionRepository interface {
	Create(ctx context.Context, amount int) (*domains.PaymentTransaction, error)
	Capture(ctx context.Context, id int) (*domains.PaymentTransaction, error)
	Cahnge(ctx context.Context, transaction domains.PaymentTransaction) (*domains.PaymentTransaction, error)
}
