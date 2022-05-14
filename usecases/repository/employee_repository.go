package repository

import (
	"clean-architecture/adapters/gateways/stores"
	"clean-architecture/domains"
	"golang.org/x/net/context"
)

type EmployeeRepository interface {
	FindEmployee(ctx context.Context, id int) (*domains.Employee, error)
	Create(ctx context.Context, e domains.Employee) (*domains.Employee, error)
}

type employeeRepository struct {
	employeeGateway stores.EmployeeGateway
}

func NewEmployeeRepository(employeeGateway stores.EmployeeGateway) EmployeeRepository {
	return &employeeRepository{employeeGateway: employeeGateway}
}

func (e employeeRepository) FindEmployee(ctx context.Context, id int) (*domains.Employee, error) {
	employee, err := e.employeeGateway.FindEmployeeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (e employeeRepository) Create(ctx context.Context, employee domains.Employee) (*domains.Employee, error) {
	result, err := e.employeeGateway.Create(ctx, employee)
	if err != nil {
		return nil, err
	}
	return result, nil
}
