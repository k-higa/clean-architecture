package mock

import (
	"clean-architecture/adapters/gateways/repository"
	"clean-architecture/usecases/test/mock"

	"clean-architecture/usecases"
	"clean-architecture/usecases/input_port"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := usecases.NewEmployeeUseCase(mock.MockRepository{}, repository.NewTransactionManger(nil))
	input := input_port.Emoployee{ID: 1}
	employee, err := usecase.FindEmployee(input)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, employee.ID, 1)
	assert.Equal(t, employee.Name, "Sam")
	assert.Equal(t, employee.Age, 29)
}
