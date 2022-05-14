package repository

import (
	"clean-architecture/adapters/gateways/clients"
	"clean-architecture/adapters/gateways/rdb"
	"clean-architecture/domains"
)

type PaymentTransactionRepository struct {
	apiKeyGateway   rdb.APIKeyGateway
	amazonPayClient clients.AmazonPayClient
}

func NewPaymentTransactionRepository(
	apiKeyGateway rdb.APIKeyGateway,
	amazonPayClient clients.AmazonPayClient,
) domains.PaymentTransactionRepository {
	return &PaymentTransactionRepository{apiKeyGateway: apiKeyGateway, amazonPayClient: amazonPayClient}
}

func (p PaymentTransactionRepository) Create(amount int) (*domains.PaymentTransaction, error) {
	apiKey, err := p.apiKeyGateway.FindByID(1)
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

func (p PaymentTransactionRepository) Capture(id int) (*domains.PaymentTransaction, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentTransactionRepository) Cahnge(transaction domains.PaymentTransaction) (*domains.PaymentTransaction, error) {
	//TODO implement me
	panic("implement me")
}
