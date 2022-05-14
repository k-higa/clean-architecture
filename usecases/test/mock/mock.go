package mock

import (
	"clean-architecture/domains"
	"golang.org/x/net/context"
)

type MockRepository struct {
}

func (m MockRepository) FindEmployee(ctx context.Context, id int) (*domains.Employee, error) {
	return &domains.Employee{
		ID:   1,
		Name: "Sam",
		Age:  29,
	}, nil
}

func (m MockRepository) Create(ctx context.Context, e domains.Employee) (*domains.Employee, error) {
	return nil, nil
}
