package rdb

import (
	"clean-architecture/domains"
	ei "clean-architecture/external_interfaces"
	"database/sql"
)

type EmployeeGateway interface {
	FindEmployeeOnly(id int) (*domains.Employee, error)
	Create(employee domains.Employee) (*domains.Employee, error)
}
type employeeGateway struct {
}

func NewEmployeeGateway() EmployeeGateway {
	return &employeeGateway{}
}

func (e employeeGateway) FindEmployeeOnly(id int) (*domains.Employee, error) {
	var employee = ei.Employee{}
	res := ei.DB.Model(ei.Employee{}).Where("id = ?", id).Find(&employee)
	if res.Error != nil {
		return nil, res.Error
	}
	return &domains.Employee{
		ID:   employee.ID,
		Name: employee.Name.String,
		Age:  int(employee.Age.Int32),
	}, nil
}

func (e employeeGateway) Create(employee domains.Employee) (*domains.Employee, error) {
	model := &ei.Employee{
		Name: sql.NullString{String: employee.Name, Valid: true},
		Age:  sql.NullInt32{Int32: int32(employee.Age), Valid: true},
	}
	res := ei.DB.Create(model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &domains.Employee{
		ID:   model.ID,
		Name: model.Name.String,
		Age:  int(model.Age.Int32),
	}, nil
}
