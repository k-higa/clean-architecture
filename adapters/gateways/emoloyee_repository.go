package gateways

import (
	"clean-architecture/database"
	"clean-architecture/entities"
	"database/sql"
)

type EmoloyeeRepository struct {
}

func NewEmoloyeeRepository() entities.EmployeeRepository {
	return &EmoloyeeRepository{}
}

func (e EmoloyeeRepository) FindEmoloyeeOnly(id int) (*entities.Emoloyee, error) {
	var emoloyee = database.Emoloyee{}
	res := database.DB.Model(database.Emoloyee{}).Where("id = ?", id).Find(&emoloyee)
	if res.Error != nil {
		return nil, res.Error
	}
	return &entities.Emoloyee{
		ID:   emoloyee.ID,
		Name: emoloyee.Name.String,
		Age:  int(emoloyee.Age.Int32),
	}, nil
}

func (e EmoloyeeRepository) Create(emoloyee entities.Emoloyee) (*entities.Emoloyee, error) {
	model := &database.Emoloyee{
		ID:   emoloyee.ID,
		Name: sql.NullString{String: emoloyee.Name, Valid: true},
		Age:  sql.NullInt32{Int32: int32(emoloyee.Age), Valid: true},
	}
	res := database.DB.Create(model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &entities.Emoloyee{
		ID:   model.ID,
		Name: model.Name.String,
		Age:  int(model.Age.Int32),
	}, nil
}
