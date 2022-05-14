package repository

import (
	"clean-architecture/adapters/gateways/clients"
	"clean-architecture/adapters/gateways/stores"
	"clean-architecture/domains"
	ei "clean-architecture/external_interfaces"
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
	amazonPayClient           clients.AmazonPayClient
}

func NewPaymentTransactionRepository(
	paymentTransactionGateway stores.PaymentTransactionGateway,
	apiKeyGateway stores.APIKeyGateway,
	amazonPayClient clients.AmazonPayClient) PaymentTransactionRepository {
	return &paymentTransactionRepository{
		paymentTransactionGateway: paymentTransactionGateway,
		apiKeyGateway:             apiKeyGateway,
		amazonPayClient:           amazonPayClient}
}

func (p paymentTransactionRepository) Entry(ctx context.Context, amount int) (*domains.PaymentTransaction, *othter.PaymentLog, error) {
	apiKey, err := p.apiKeyGateway.FindByID(ctx, 1)
	if err != nil {
		return nil, nil, err
	}
	req := clients.AmazonPayReq{
		Amount: amount,
	}
	res, err := p.amazonPayClient.Entry(req, apiKey.KEY.String)
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
