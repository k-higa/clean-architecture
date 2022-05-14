package rdb

import (
	ei "clean-architecture/external_interfaces"
)

type APIKeyGateway interface {
	FindByID(id int) (*ei.APIKey, error)
}
type apiKeyGateway struct {
}

func NewAPIKeyGateway() APIKeyGateway {
	return &apiKeyGateway{}
}

func (a apiKeyGateway) FindByID(id int) (*ei.APIKey, error) {
	var apiKey = ei.APIKey{}
	res := ei.DB.Model(ei.APIKey{}).Where("id = ?", id).Find(&apiKey)
	if res.Error != nil {
		return nil, res.Error
	}
	return &apiKey, nil
}
