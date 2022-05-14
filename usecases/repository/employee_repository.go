package repository

import (
	"clean-architecture/domains"
	"golang.org/x/net/context"
)

type EmployeeRepository interface {
	FindEmployee(ctx context.Context, id int) (*domains.Employee, error)
	Create(ctx context.Context, e domains.Employee) (*domains.Employee, error)
}
