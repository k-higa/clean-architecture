package repository

import (
	"clean-architecture/adapters/gateways/stores"
	"clean-architecture/domains"
	"clean-architecture/usecases"
)

type EmployeeRepository interface {
	FindEmployee(ctx usecases.Context, id int) (*domains.Employee, error)
	Create(ctx usecases.Context, e domains.Employee) (*domains.Employee, error)
}

type employeeRepository struct {
	employeeGateway stores.EmployeeGateway
}

func NewEmployeeRepository(employeeGateway stores.EmployeeGateway) EmployeeRepository {
	return &employeeRepository{employeeGateway: employeeGateway}
}

func (e employeeRepository) FindEmployee(ctx usecases.Context, id int) (*domains.Employee, error) {
	employee, err := e.employeeGateway.FindEmployeeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (e employeeRepository) Create(ctx usecases.Context, employee domains.Employee) (*domains.Employee, error) {
	result, err := e.employeeGateway.Create(ctx, employee)
	if err != nil {
		return nil, err
	}
	return result, nil
}
