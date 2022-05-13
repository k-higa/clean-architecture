package usecases

import (
	"clean-architecture/domains"
	"clean-architecture/usecases/input_port"
	"clean-architecture/usecases/output_port"
)

type EmployeeUseCase interface {
	FindEmployee(d input_port.Emoployee) (*output_port.Emoployee, error)
}

type EmployeeInteractor struct {
	employeeRepo domains.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo domains.EmployeeRepository) EmployeeUseCase {
	return &EmployeeInteractor{employeeRepo: employeeRepo}
}

func (e EmployeeInteractor) FindEmployee(d input_port.Emoployee) (*output_port.Emoployee, error) {
	employee, err := e.employeeRepo.FindEmoloyeeOnly(d.ID)
	if err != nil {
		return nil, err
	}
	out := &output_port.Emoployee{
		ID:   employee.ID,
		Name: employee.Name,
		Age:  employee.Age,
	}
	return out, nil
}
