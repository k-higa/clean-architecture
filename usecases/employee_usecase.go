package usecases

import (
	"clean-architecture/domains"
	"clean-architecture/usecases/input_port"
	"clean-architecture/usecases/output_port"
	"clean-architecture/usecases/repository"
	"golang.org/x/net/context"
)

type EmployeeUseCase interface {
	FindEmployee(d input_port.Emoployee) (*output_port.Emoployee, error)
	CreatedEmployee(d input_port.Emoployee) (*output_port.Emoployee, error)
}

type EmployeeInteractor struct {
	employeeRepo repository.EmployeeRepository
	tm           repository.TransactionManager
}

func NewEmployeeUseCase(
	employeeRepo repository.EmployeeRepository,
	tm repository.TransactionManager) EmployeeUseCase {
	return &EmployeeInteractor{employeeRepo: employeeRepo, tm: tm}
}

func (e EmployeeInteractor) FindEmployee(d input_port.Emoployee) (*output_port.Emoployee, error) {
	employee, err := e.employeeRepo.FindEmployee(context.TODO(), d.ID)
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
	emoloyee := domains.Employee{
		Name: d.Name,
		Age:  d.Age,
	}
	var out output_port.Emoployee
	err := e.tm.Transact(func(ctx context.Context) error {
		//transaction
		res, err := e.employeeRepo.Create(ctx, emoloyee)
		if err != nil {
			return err
		}
		res, err = e.employeeRepo.Create(ctx, emoloyee)
		if err != nil {
			return err
		}
		out.ID = res.ID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &out, nil
}
