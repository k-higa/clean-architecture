package stores

import (
	"clean-architecture/adapters/gateways/stores/models"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type PaymentTransactionGateway interface {
	Save(ctx context.Context, pt models.PaymentTransaction) (*models.PaymentTransaction, error)
}
type paymentTransactionGateway struct {
}

func NewPaymentTransactionGateway() PaymentTransactionGateway {
	return &paymentTransactionGateway{}
}

func (p paymentTransactionGateway) Save(ctx context.Context, pt models.PaymentTransaction) (*models.PaymentTransaction, error) {
	tx := ctx.Value("TX").(gorm.DB)
	res := tx.Model(models.PaymentTransaction{}).Save(pt)
	if res.Error != nil {
		return nil, res.Error
	}
	return &pt, nil
}
