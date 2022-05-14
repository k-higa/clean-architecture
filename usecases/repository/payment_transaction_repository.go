package repository

import (
	"clean-architecture/adapters/gateways/api"
	"clean-architecture/adapters/gateways/api/input"
	"clean-architecture/adapters/gateways/stores"
	ei "clean-architecture/adapters/gateways/stores/models"
	"clean-architecture/domains"
	"clean-architecture/othter"
	"golang.org/x/net/context"
)

type PaymentTransactionRepository interface {
	Entry(ctx context.Context, amount int) (*domains.PaymentTransaction, *othter.PaymentLog, error)
	Capture(ctx context.Context, id int) (*domains.PaymentTransaction, *othter.PaymentLog, error)
	Cahnge(ctx context.Context, transaction domains.PaymentTransaction) (*domains.PaymentTransaction, *othter.PaymentLog, error)
}

type paymentTransactionRepository struct {
	paymentTransactionGateway stores.PaymentTransactionGateway
	apiKeyGateway             stores.APIKeyGateway
	amazonPayAPIGateway       api.AmazonPayAPIGateway
}

func NewPaymentTransactionRepository(
	paymentTransactionGateway stores.PaymentTransactionGateway,
	apiKeyGateway stores.APIKeyGateway,
	amazonPayClient api.AmazonPayAPIGateway) PaymentTransactionRepository {
	return &paymentTransactionRepository{
		paymentTransactionGateway: paymentTransactionGateway,
		apiKeyGateway:             apiKeyGateway,
		amazonPayAPIGateway:       amazonPayClient}
}

func (p paymentTransactionRepository) Entry(ctx context.Context, amount int) (*domains.PaymentTransaction, *othter.PaymentLog, error) {
	apiKey, err := p.apiKeyGateway.FindByID(ctx, 1)
	if err != nil {
		return nil, nil, err
	}
	req := input.AmazonPayReq{
		Amount: amount,
	}
	res, err := p.amazonPayAPIGateway.Entry(req, apiKey.KEY.String)
	if err != nil {
		return nil, nil, err
	}
	model := ei.PaymentTransaction{
		ID:     res.ID,
		Amount: res.Amount,
		Status: res.Status,
	}
	log := &othter.PaymentLog{
		Host: res.Host, ContentType: res.ContentType,
		Request: res.RequestLog, Response: res.ResponseLog,
	}
	pt, err := p.paymentTransactionGateway.Save(ctx, model)
	if err != nil {
		return nil, log, err
	}
	return &domains.PaymentTransaction{
		ID:     pt.ID,
		Amount: pt.Amount,
		Status: pt.Status,
	}, log, nil
}

func (p paymentTransactionRepository) Capture(ctx context.Context, id int) (*domains.PaymentTransaction, *othter.PaymentLog, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentTransactionRepository) Cahnge(ctx context.Context, transaction domains.PaymentTransaction) (*domains.PaymentTransaction, *othter.PaymentLog, error) {
	//TODO implement me
	panic("implement me")
}
