package repository

import (
	"clean-architecture/adapters/gateways/clients"
	"clean-architecture/adapters/gateways/stores"
	"clean-architecture/domains"
	"clean-architecture/usecases/repository"
	"golang.org/x/net/context"
)

type employeeRepository struct {
	employeeGateway   stores.EmployeeGateway
	employeeAPIClient clients.EmployeeClient
}

func NewEmployeeRepository(employeeGateway stores.EmployeeGateway) repository.EmployeeRepository {
	return &employeeRepository{employeeGateway: employeeGateway}
}

func (e employeeRepository) FindEmployee(ctx context.Context, id int) (*domains.Employee, error) {
	employee, err := e.employeeGateway.FindEmployeeOnly(ctx, id) //stores version
	employee, err = e.employeeAPIClient.Fetch(id)                //external api version
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
