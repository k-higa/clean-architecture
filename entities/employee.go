package entities

/**
ドメインエンティティ
*/
type Employee struct {
	ID   int
	Name string
	Age  int
}

type EmployeeRepository interface {
	FindEmployeeOnly(id int) (*Employee, error)
	Create(e Employee) (*Employee, error)
}
