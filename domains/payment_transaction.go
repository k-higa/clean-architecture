package domains

type PaymentTransaction struct {
	ID     int
	Amount int
	Status string
}

type PaymentTransactionRepository interface {
	Create(amount int) (*PaymentTransaction, error)
	Capture(id int) (*PaymentTransaction, error)
	Cahnge(transaction PaymentTransaction) (*PaymentTransaction, error)
}
