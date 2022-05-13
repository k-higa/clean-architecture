package repositories

type EmoloyeeClient interface {
	Fetch(req EmployeeReq) (*EmployeeRes, error)
}
type emoloyeeClient struct {
}

func NewEmoloyeeClient() EmoloyeeClient {
	return &emoloyeeClient{}
}

func (e *emoloyeeClient) Fetch(req EmployeeReq) (*EmployeeRes, error) {
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
