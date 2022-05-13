package mock

import (
	"clean-architecture/domains"
	"clean-architecture/usecases"
	"clean-architecture/usecases/input_port"
	"clean-architecture/usecases/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックの生成
	mockRepo := mock.NewMockEmployeeRepository(ctrl)
	out := &domains.Emoloyee{
		ID:   1,
		Name: "Sam",
		Age:  29,
	}
	mockRepo.EXPECT().FindEmoloyeeOnly(gomock.Any()).Return(out, nil)
	usecase := usecases.NewEmployeeUsecase(mockRepo)
	input := input_port.Emoployee{ID: 1}
	employee, err := usecase.FindEmployee(input)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, employee.ID, 1)
	assert.Equal(t, employee.Name, "Sam")
	assert.Equal(t, employee.Age, 29)
}
