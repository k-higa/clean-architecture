package repositories

import (
	"clean-architecture/database"
	"clean-architecture/entities"
	"database/sql"
)

type EmployeeRepository struct {
}

func NewEmployeeRepository() entities.EmployeeRepository {
	return &EmployeeRepository{}
}

func (e EmployeeRepository) FindEmployeeOnly(id int) (*entities.Employee, error) {
	var employee = database.Emoloyee{}
	res := database.DB.Model(database.Emoloyee{}).Where("id = ?", id).Find(&employee)
	if res.Error != nil {
		return nil, res.Error
	}
	return &entities.Employee{
		ID:   employee.ID,
		Name: employee.Name.String,
		Age:  int(employee.Age.Int32),
	}, nil
}

func (e EmployeeRepository) Create(employee entities.Employee) (*entities.Employee, error) {
	model := &database.Emoloyee{
		ID:   employee.ID,
		Name: sql.NullString{String: employee.Name, Valid: true},
		Age:  sql.NullInt32{Int32: int32(employee.Age), Valid: true},
	}
	res := database.DB.Create(model)
	if res.Error != nil {
		return nil, res.Error
	}
	return &entities.Employee{
		ID:   model.ID,
		Name: model.Name.String,
		Age:  int(model.Age.Int32),
	}, nil
}
