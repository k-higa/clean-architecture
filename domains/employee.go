package domains

/**
Domain entiry
*/
type Employee struct {
	ID   int
	Name string
	Age  int
}

type EmployeeRepository interface {
	FindEmployee(id int) (*Employee, error)
	Create(e Employee) (*Employee, error)
}
