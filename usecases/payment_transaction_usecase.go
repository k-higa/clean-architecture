package usecases

import (
	"clean-architecture/domains"
	"clean-architecture/usecases/input_port"
	"clean-architecture/usecases/output_port"
)

type PaymentTransactionUseCase interface {
	Entry(d input_port.PaymentEntry) (*output_port.PaymentTransaction, error)
}

type paymentTransactionInteractor struct {
	paymentTransactionRepo domains.PaymentTransactionRepository
}

func NewPaymentTransactionUseCase(paymentTransactionRepo domains.PaymentTransactionRepository) PaymentTransactionUseCase {
	return &paymentTransactionInteractor{paymentTransactionRepo: paymentTransactionRepo}
}

func (p paymentTransactionInteractor) Entry(d input_port.PaymentEntry) (*output_port.PaymentTransaction, error) {
	transaction, err := p.paymentTransactionRepo.Create(d.Amount)
	if err != nil {
		return nil, err
	}
	out := &output_port.PaymentTransaction{
		ID:     transaction.ID,
		Amount: transaction.Amount,
		Status: transaction.Status,
	}
	return out, nil
}
