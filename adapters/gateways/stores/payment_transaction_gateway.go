package stores

import (
	ei "clean-architecture/external_interfaces"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type PaymentTransactionGateway interface {
	Save(ctx context.Context, pt ei.PaymentTransaction) (*ei.PaymentTransaction, error)
}
type paymentTransactionGateway struct {
}

func NewPaymentTransactionGateway() PaymentTransactionGateway {
	return &paymentTransactionGateway{}
}

func (p paymentTransactionGateway) Save(ctx context.Context, pt ei.PaymentTransaction) (*ei.PaymentTransaction, error) {
	tx := ctx.Value("TX").(gorm.DB)
	res := tx.Model(ei.PaymentTransaction{}).Save(pt)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pt, nil
}
