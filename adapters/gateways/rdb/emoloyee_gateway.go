package rdb

import (
	"clean-architecture/domains"
	ei "clean-architecture/external_interfaces"
	"database/sql"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type EmployeeGateway interface {
	FindEmployeeOnly(ctx context.Context, id int) (*domains.Employee, error)
	Create(ctx context.Context, employee domains.Employee) (*domains.Employee, error)
}
type employeeGateway struct {
}

func NewEmployeeGateway() EmployeeGateway {
	return &employeeGateway{}
}

func (e employeeGateway) FindEmployeeOnly(ctx context.Context, id int) (*domains.Employee, error) {
	var employee = ei.Employee{}
	tx := ctx.Value("TX").(gorm.DB)
	res := tx.Model(ei.Employee{}).Where("id = ?", id).Find(&employee)
	if res.Error != nil {
		return nil, res.Error
	}
	return &domains.Employee{
		ID:   employee.ID,
		Name: employee.Name.String,
		Age:  int(employee.Age.Int32),
	}, nil
}

func (e employeeGateway) Create(ctx context.Context, employee domains.Employee) (*domains.Employee, error) {
	model := &ei.Employee{
		Name: sql.NullString{String: employee.Name, Valid: true},
		Age:  sql.NullInt32{Int32: int32(employee.Age), Valid: true},
	}
	tx := ctx.Value("TX").(gorm.DB)
	res := tx.Create(model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &domains.Employee{
		ID:   model.ID,
		Name: model.Name.String,
		Age:  int(model.Age.Int32),
	}, nil
}
