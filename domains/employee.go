package domains

/**
ドメインエンティティ
*/
type Emoloyee struct {
	ID   int
	Name string
	Age  int
}

type EmployeeRepository interface {
	FindEmoloyeeOnly(id int) (*Emoloyee, error)
	Create(e Emoloyee) (*Emoloyee, error)
}
