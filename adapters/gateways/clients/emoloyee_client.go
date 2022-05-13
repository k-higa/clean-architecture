package repositories

type EmployeeClient interface {
	Fetch(req EmployeeReq) (*EmployeeRes, error)
}
type employeeClient struct {
}

func NewEmployeeClient() EmployeeClient {
	return &employeeClient{}
}

func (e *employeeClient) Fetch(req EmployeeReq) (*EmployeeRes, error) {
	return nil, nil
}

type EmployeeReq struct {
	ID int
}

type EmployeeRes struct {
	ID   int
	Name string
	Age  int
}
