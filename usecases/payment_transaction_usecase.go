package usecases

import (
	"clean-architecture/usecases/input_port"
	"clean-architecture/usecases/output_port"
	"clean-architecture/usecases/repository"
)

type PaymentTransactionUseCase interface {
	Entry(d input_port.PaymentEntry) (*output_port.PaymentTransaction, error)
}

type paymentTransactionInteractor struct {
	paymentTransactionRepo repository.PaymentTransactionRepository
	paymentLogReoisitory   repository.PaymentLogReoisitory
	tm                     repository.TransactionManager
}

func NewPaymentTransactionUseCase(
	paymentTransactionRepo repository.PaymentTransactionRepository,
	paymentLogReoisitory repository.PaymentLogReoisitory,
	tm repository.TransactionManager) PaymentTransactionUseCase {
	return &paymentTransactionInteractor{
		paymentTransactionRepo: paymentTransactionRepo,
		paymentLogReoisitory:   paymentLogReoisitory,
		tm:                     tm}
}

func (p paymentTransactionInteractor) Entry(d input_port.PaymentEntry) (*output_port.PaymentTransaction, error) {
	res, log, err := p.paymentTransactionRepo.Entry(p.tm.Tx(), d.Amount)
	if err != nil {
		return nil, err
	}
	err = p.paymentLogReoisitory.Create(p.tm.Tx(), *log)
	if err != nil {
		return nil, err
	}
	out := &output_port.PaymentTransaction{
		ID:     res.ID,
		Amount: res.Amount,
		Status: res.Status,
	}
	return out, nil
}
