package api

import (
	"clean-architecture/adapters/gateways/api/input"
	"clean-architecture/adapters/gateways/api/output"
)

type AmazonPayAPIGateway interface {
	Entry(req input.AmazonPayReq, apiKey string) (*output.AmazonPayRes, error)
}

type amazonPayAPIGateway struct {
}

func NewAmazonPayAPIGateway() AmazonPayAPIGateway {
	return &amazonPayAPIGateway{}
}

func (a *amazonPayAPIGateway) Entry(req input.AmazonPayReq, apiKey string) (*output.AmazonPayRes, error) {
	return &output.AmazonPayRes{ID: 1, Amount: 100, Status: "authorize"}, nil
}
