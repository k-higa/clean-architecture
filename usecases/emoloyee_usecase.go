package usecases

import (
	"clean-architecture/entities"
	"clean-architecture/usecases/input_port"
	"clean-architecture/usecases/output_port"
)

type EmployeeUseCase interface {
	FindEmployee(d input_port.Emoployee) (*output_port.Emoployee, error)
	CreatedEmployee(d input_port.Emoployee) (*output_port.Emoployee, error)
}

type EmployeeInteractor struct {
	employeeRepo entities.EmployeeRepository
}

func NewEmployeeUsecase(employeeRepo entities.EmployeeRepository) EmployeeUseCase {
	return &EmployeeInteractor{employeeRepo: employeeRepo}
}

func (e EmployeeInteractor) FindEmployee(d input_port.Emoployee) (*output_port.Emoployee, error) {
	employee, err := e.employeeRepo.FindEmployeeOnly(d.ID)
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

func (e EmployeeInteractor) CreatedEmployee(d input_port.Emoployee) (*output_port.Emoployee, error) {
	emoloyee := entities.Employee{
		Name: d.Name,
		Age:  d.Age,
	}
	employee, err := e.employeeRepo.Create(emoloyee)
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
