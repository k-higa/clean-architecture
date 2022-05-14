package clients

import (
	"clean-architecture/domains"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type EmployeeClient interface {
	Fetch(id int) (*domains.Employee, error)
}
type employeeClient struct {
}

func NewEmployeeClient() EmployeeClient {
	return &employeeClient{}
}

func (e *employeeClient) Fetch(id int) (*domains.Employee, error) {
	body, _ := json.Marshal(EmployeeReq{ID: id})
	res, err := http.Post("localhost", "application/json", strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	var result EmployeeRes
	if err = json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return &domains.Employee{
		ID:   result.ID,
		Name: result.Name,
		Age:  result.Age,
	}, nil
}

type EmployeeReq struct {
	ID int
}

type EmployeeRes struct {
	ID   int
	Name string
	Age  int
}
