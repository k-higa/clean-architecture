package stores

import (
	"clean-architecture/adapters/gateways/stores/models"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type APIKeyGateway interface {
	FindByID(ctx context.Context, id int) (*models.APIKey, error)
}
type apiKeyGateway struct {
}

func NewAPIKeyGateway() APIKeyGateway {
	return &apiKeyGateway{}
}

func (a apiKeyGateway) FindByID(ctx context.Context, id int) (*models.APIKey, error) {
	var apiKey = models.APIKey{}
	tx := ctx.Value("TX").(gorm.DB)
	res := tx.Model(models.APIKey{}).Where("id = ?", id).Find(&apiKey)
	if res.Error != nil {
		return nil, res.Error
	}
	return &apiKey, nil
}
