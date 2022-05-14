package stores

import (
	ei "clean-architecture/external_interfaces"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type APIKeyGateway interface {
	FindByID(ctx context.Context, id int) (*ei.APIKey, error)
}
type apiKeyGateway struct {
}

func NewAPIKeyGateway() APIKeyGateway {
	return &apiKeyGateway{}
}

func (a apiKeyGateway) FindByID(ctx context.Context, id int) (*ei.APIKey, error) {
	var apiKey = ei.APIKey{}
	tx := ctx.Value("TX").(gorm.DB)
	res := tx.Model(ei.APIKey{}).Where("id = ?", id).Find(&apiKey)
	if res.Error != nil {
		return nil, res.Error
	}
	return &apiKey, nil
}
