package stores

import (
	ei "clean-architecture/external_interfaces"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type PaymentLogGateway interface {
	Save(ctx context.Context, pt ei.PaymentLog) (*ei.PaymentLog, error)
}
type paymentLogGateway struct {
}

func NewPaymentLogGateway() PaymentLogGateway {
	return &paymentLogGateway{}
}

func (p paymentLogGateway) Save(ctx context.Context, log ei.PaymentLog) (*ei.PaymentLog, error) {
	tx := ctx.Value("TX").(gorm.DB)
	res := tx.Model(ei.PaymentLog{}).Save(log)
	if res.Error != nil {
		return nil, res.Error
	}
	return &log, nil
}
