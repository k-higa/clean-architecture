package gateways

import (
	"clean-architecture/database"
	"clean-architecture/domains"
	"database/sql"
)

type EmoloyeeGateway struct {
}

func NewEmoloyeeRepository() domains.EmployeeRepository {
	return &EmoloyeeGateway{}
}

func (e EmoloyeeGateway) FindEmoloyeeOnly(id int) (*domains.Emoloyee, error) {
	var emoloyee = database.Emoloyee{}
	res := database.DB.Model(database.Emoloyee{}).Where("id = ?", id).Find(&emoloyee)
	if res.Error != nil {
		return nil, res.Error
	}
	return &domains.Emoloyee{
		ID:   emoloyee.ID,
		Name: emoloyee.Name.String,
		Age:  int(emoloyee.Age.Int32),
	}, nil
}

func (e EmoloyeeGateway) Save(emoloyee domains.Emoloyee) (*domains.Emoloyee, error) {
	model := &database.Emoloyee{
		ID:   emoloyee.ID,
		Name: sql.NullString{String: emoloyee.Name, Valid: true},
		Age:  sql.NullInt32{Int32: int32(emoloyee.Age), Valid: true},
	}
	res := database.DB.Save(model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &domains.Emoloyee{
		ID:   model.ID,
		Name: model.Name.String,
		Age:  int(model.Age.Int32),
	}, nil
}
