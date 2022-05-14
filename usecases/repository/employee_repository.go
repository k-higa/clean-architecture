package repository

import (
	"clean-architecture/adapters/gateways/clients"
	"clean-architecture/adapters/gateways/rdb"
	"clean-architecture/domains"
)

type EmployeeRepository struct {
	employeeGateway   rdb.EmployeeGateway
	employeeAPIClient clients.EmployeeClient
}

func NewEmployeeRepository(employeeGateway rdb.EmployeeGateway) domains.EmployeeRepository {
	return &EmployeeRepository{employeeGateway: employeeGateway}
}

func (e EmployeeRepository) FindEmployee(id int) (*domains.Employee, error) {
	employee, err := e.employeeGateway.FindEmployeeOnly(id) //rdb version
	employee, err = e.employeeAPIClient.Fetch(id)           //external api version
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (e EmployeeRepository) Create(employee domains.Employee) (*domains.Employee, error) {
	result, err := e.employeeGateway.Create(employee)
	if err != nil {
		return nil, err
	}
	return result, nil
}
