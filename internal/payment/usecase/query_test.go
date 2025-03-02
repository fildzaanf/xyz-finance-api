package usecase

import (
	"errors"
	"testing"
	"xyz-finance-api/internal/payment/domain"
	du "xyz-finance-api/internal/user/domain"
	"xyz-finance-api/pkg/constant"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
)

func TestGetPaymentByID(t *testing.T) {
	mockPaymentQueryRepo := new(mocks.PaymentQueryRepositoryInterface)
	mockPaymentCommandRepo := new(mocks.PaymentCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)
	paymentQueryUsecase := NewPaymentQueryUsecase(mockPaymentQueryRepo, mockPaymentCommandRepo, mockUserQueryRepo)

	t.Run("Valid GetPaymentByID", func(t *testing.T) {
		paymentID := "payment123"
		userID := "user123"
		mockPaymentQueryRepo.On("GetPaymentByID", paymentID, userID).Return(domain.Payment{ID: paymentID}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil)

		payment, err := paymentQueryUsecase.GetPaymentByID(paymentID, userID)

		mockPaymentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Equal(t, payment.ID, paymentID)
	})

	t.Run("Invalid GetPaymentByID (Empty Payment ID)", func(t *testing.T) {
		paymentID := ""
		userID := "user123"

		_, err := paymentQueryUsecase.GetPaymentByID(paymentID, userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)

		mockPaymentQueryRepo.AssertExpectations(t)
	})

	t.Run("GetPaymentByID Payment Not Found", func(t *testing.T) {
		paymentID := "nonexistent-payment"
		userID := "user123"

		mockPaymentQueryRepo.On("GetPaymentByID", paymentID, userID).Return(domain.Payment{}, errors.New(constant.ERROR_ID_NOTFOUND))

		_, err := paymentQueryUsecase.GetPaymentByID(paymentID, userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_NOTFOUND)

		mockPaymentQueryRepo.AssertExpectations(t)
	})

	t.Run("GetPaymentByID User Not Found", func(t *testing.T) {
		paymentID := "payment123"
		userID := "nonexistent-user"

		mockPaymentQueryRepo.On("GetPaymentByID", paymentID, userID).Return(domain.Payment{ID: paymentID}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New(constant.ERROR_ROLE_ACCESS))

		_, err := paymentQueryUsecase.GetPaymentByID(paymentID, userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ROLE_ACCESS)

		mockPaymentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
	})
}


func TestGetAllPayments(t *testing.T) {
	mockPaymentQueryRepo := new(mocks.PaymentQueryRepositoryInterface)
	mockPaymentCommandRepo := new(mocks.PaymentCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)
	paymentQueryUsecase := NewPaymentQueryUsecase(mockPaymentQueryRepo, mockPaymentCommandRepo, mockUserQueryRepo)

	t.Run("Valid GetAllPayments", func(t *testing.T) {
		userID := "user123"
		mockPaymentQueryRepo.On("GetAllPayments", userID).Return([]domain.Payment{{ID: "payment123"}}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil)

		payments, err := paymentQueryUsecase.GetAllPayments(userID)

		mockPaymentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		assert.Nil(t, err)
		assert.Len(t, payments, 1)
		assert.Equal(t, payments[0].ID, "payment123")
	})

	t.Run("Invalid GetAllPayments (Empty User ID)", func(t *testing.T) {
		userID := ""

		_, err := paymentQueryUsecase.GetAllPayments(userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ID_INVALID)

		mockPaymentQueryRepo.AssertExpectations(t)
	})

	t.Run("GetAllPayments Data Not Found", func(t *testing.T) {
		userID := "user123"
	
		mockPaymentQueryRepo.On("GetAllPayments", userID).Return(nil, errors.New(constant.ERROR_DATA_EMPTY)).Once()
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{ID: userID}, nil).Once()
	

		payments, err := paymentQueryUsecase.GetAllPayments(userID)
	
		assert.Nil(t, payments)
		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_DATA_EMPTY)
	
		mockPaymentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
	})
	
	

	t.Run("GetAllPayments User Not Found", func(t *testing.T) {
		userID := "nonexistent-user"

		mockPaymentQueryRepo.On("GetAllPayments", userID).Return([]domain.Payment{}, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New(constant.ERROR_ROLE_ACCESS))

		_, err := paymentQueryUsecase.GetAllPayments(userID)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_ROLE_ACCESS)

		mockPaymentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
	})
}
