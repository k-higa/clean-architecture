package stores

import (
	"clean-architecture/adapters/gateways/stores/models"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type PaymentLogGateway interface {
	Save(ctx context.Context, pt models.PaymentLog) (*models.PaymentLog, error)
}
type paymentLogGateway struct {
}

func NewPaymentLogGateway() PaymentLogGateway {
	return &paymentLogGateway{}
}

func (p paymentLogGateway) Save(ctx context.Context, log models.PaymentLog) (*models.PaymentLog, error) {
	tx := ctx.Value("TX").(gorm.DB)
	res := tx.Model(models.PaymentLog{}).Save(log)
	if res.Error != nil {
		return nil, res.Error
	}
	return &log, nil
}
