package repository

import (
	"clean-architecture/adapters/gateways/stores"
	ei "clean-architecture/adapters/gateways/stores/models"
	"clean-architecture/othter"
	"clean-architecture/usecases"
	"database/sql"
)

type PaymentLogReoisitory interface {
	Create(ctx usecases.Context, pl othter.PaymentLog) error
}

type PaymentLog struct {
}

type paymentLogReoisitory struct {
	paymentLogGateway stores.PaymentLogGateway
}

func NewPaymentLogReoisitory(paymentLogGateway stores.PaymentLogGateway) *paymentLogReoisitory {
	return &paymentLogReoisitory{paymentLogGateway: paymentLogGateway}
}

func (p paymentLogReoisitory) Create(ctx usecases.Context, pi othter.PaymentLog) error {
	log := ei.PaymentLog{
		Host:        sql.NullString{String: pi.Host, Valid: true},
		ContentType: sql.NullString{String: pi.Host, Valid: true},
		Request:     sql.NullString{String: pi.Host, Valid: true},
		Response:    sql.NullString{String: pi.Host, Valid: true},
	}
	_, err := p.paymentLogGateway.Save(ctx, log)
	return err
}
