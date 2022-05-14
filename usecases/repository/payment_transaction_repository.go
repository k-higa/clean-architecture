package repository

import (
	"clean-architecture/adapters/gateways/clients"
	"clean-architecture/adapters/gateways/stores"
	"clean-architecture/domains"
	"golang.org/x/net/context"
)

type PaymentTransactionRepository interface {
	Create(ctx context.Context, amount int) (*domains.PaymentTransaction, error)
	Capture(ctx context.Context, id int) (*domains.PaymentTransaction, error)
	Cahnge(ctx context.Context, transaction domains.PaymentTransaction) (*domains.PaymentTransaction, error)
}

type paymentTransactionRepository struct {
	apiKeyGateway   stores.APIKeyGateway
	amazonPayClient clients.AmazonPayClient
}

func NewPaymentTransactionRepository(
	apiKeyGateway stores.APIKeyGateway,
	amazonPayClient clients.AmazonPayClient) PaymentTransactionRepository {
	return &paymentTransactionRepository{apiKeyGateway: apiKeyGateway, amazonPayClient: amazonPayClient}
}

func (p paymentTransactionRepository) Create(ctx context.Context, amount int) (*domains.PaymentTransaction, error) {
	apiKey, err := p.apiKeyGateway.FindByID(ctx, 1)
	if err != nil {
		return nil, err
	}
	req := clients.AmazonPayReq{
		Amount: amount,
	}
	res, err := p.amazonPayClient.Entry(req, apiKey.KEY.String)
	if err != nil {
		return nil, err
	}
	return &domains.PaymentTransaction{
		ID:     res.ID,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}

func (p paymentTransactionRepository) Capture(ctx context.Context, id int) (*domains.PaymentTransaction, error) {
	//TODO implement me
	panic("implement me")
}

func (p paymentTransactionRepository) Cahnge(ctx context.Context, transaction domains.PaymentTransaction) (*domains.PaymentTransaction, error) {
	//TODO implement me
	panic("implement me")
}
