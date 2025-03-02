package usecase

import (
	"errors"
	"testing"
	"time"
	di "xyz-finance-api/internal/installment/domain"
	"xyz-finance-api/internal/payment/domain"
	du "xyz-finance-api/internal/user/domain"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePayment(t *testing.T) {
	mockPaymentCommandRepo := new(mocks.PaymentCommandRepositoryInterface)
	mockPaymentQueryRepo := new(mocks.PaymentQueryRepositoryInterface)
	mockInstallmentQueryRepo := new(mocks.InstallmentQueryRepositoryInterface)
	mockInstallmentCommandRepo := new(mocks.InstallmentCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)

	paymentUsecase := NewPaymentCommandUsecase(
		mockPaymentCommandRepo,
		mockPaymentQueryRepo,
		mockInstallmentQueryRepo,
		mockInstallmentCommandRepo,
		nil,
		mockUserQueryRepo,
	)

	t.Run("Valid Payment", func(t *testing.T) {
		payment := domain.Payment{
			InstallmentID: "installment123",
		}
		userID := "user123"
		installment := di.Installment{
			ID:     "installment123",
			Amount: 1000,
		}
		user := du.User{
			ID: "user123",
		}

		expectedPayment := domain.Payment{
			InstallmentID: "installment123",
			GrossAmount:   installment.Amount,
			Status:        "pending",
			PaymentURL:    "https://app.sandbox.midtrans.com/snap/v4/redirection/6f475f4c-18cb-446c-8bad-b484b7cf9773",
			Token:         "6f475f4c-18cb-446c-8bad-b484b7cf9773",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		mockInstallmentQueryRepo.On("GetInstallmentByID", payment.InstallmentID, userID).Return(installment, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(user, nil)
		mockPaymentCommandRepo.On("CreatePayment", mock.MatchedBy(func(p domain.Payment) bool {
			return p.InstallmentID == expectedPayment.InstallmentID &&
				p.GrossAmount == expectedPayment.GrossAmount &&
				p.PaymentURL != "" && p.Token != "" &&
				p.Status == "pending"
		})).Return(expectedPayment, nil)

		result, err := paymentUsecase.CreatePayment(payment, userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedPayment.InstallmentID, result.InstallmentID)
		assert.Equal(t, expectedPayment.GrossAmount, result.GrossAmount)
		assert.Equal(t, expectedPayment.PaymentURL, result.PaymentURL)
		assert.Equal(t, expectedPayment.Token, result.Token)
		assert.Equal(t, "pending", result.Status)

		mockInstallmentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
		mockPaymentCommandRepo.AssertExpectations(t)
	})

	t.Run("Installment Not Found", func(t *testing.T) {
		payment := domain.Payment{
			InstallmentID: "installment123",
		}
		userID := "user123"

		mockInstallmentQueryRepo.On("GetInstallmentByID", payment.InstallmentID, userID).Return(di.Installment{}, errors.New("installment not found"))

		result, err := paymentUsecase.CreatePayment(payment, userID)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "installment not found")
		assert.Equal(t, result, domain.Payment{})
		mockInstallmentQueryRepo.AssertExpectations(t)
	})

	t.Run("User Not Found", func(t *testing.T) {
		payment := domain.Payment{
			InstallmentID: "installment123",
		}
		userID := "user123"
		installment := di.Installment{
			ID:     "installment123",
			Amount: 1000,
		}

		mockInstallmentQueryRepo.On("GetInstallmentByID", payment.InstallmentID, userID).Return(installment, nil)
		mockUserQueryRepo.On("GetUserByID", userID).Return(du.User{}, errors.New("user not found"))

		result, err := paymentUsecase.CreatePayment(payment, userID)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "user not found")
		assert.Equal(t, result, domain.Payment{})
		mockInstallmentQueryRepo.AssertExpectations(t)
		mockUserQueryRepo.AssertExpectations(t)
	})

	t.Run("Create Payment with Empty InstallmentID", func(t *testing.T) {
		userID := "user123"
		payment := domain.Payment{
			InstallmentID: "", 
		}
	
		
		_, err := paymentUsecase.CreatePayment(payment, userID)

		assert.Error(t, err)
		mockInstallmentCommandRepo.AssertExpectations(t)
	})
	
	
}

func TestUpdatePaymentStatus(t *testing.T) {
	mockPaymentCommandRepo := new(mocks.PaymentCommandRepositoryInterface)
	mockPaymentQueryRepo := new(mocks.PaymentQueryRepositoryInterface)
	mockInstallmentQueryRepo := new(mocks.InstallmentQueryRepositoryInterface)
	mockInstallmentCommandRepo := new(mocks.InstallmentCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)

	paymentUsecase := NewPaymentCommandUsecase(
		mockPaymentCommandRepo,
		mockPaymentQueryRepo,
		mockInstallmentQueryRepo,
		mockInstallmentCommandRepo,
		nil,
		mockUserQueryRepo,
	)
	t.Run("Valid Payment Status Update", func(t *testing.T) {
		payment := domain.Payment{
			InstallmentID: "installment123",
			Status:        "pending",
		}

		installmentPaid := di.Installment{
			Status: "paid",
		}
	
		mockPaymentQueryRepo.On("GetPaymentByInstallmentID", payment.InstallmentID).Return(payment, nil)
		
		mockPaymentQueryRepo.On("UpdatePaymentStatus", payment.InstallmentID, "settlement").Return(nil)
	
		mockInstallmentCommandRepo.On("UpdateInstallmentStatusByID", payment.InstallmentID, installmentPaid).Return(installmentPaid, nil)
	
		err := paymentUsecase.UpdatePaymentStatus(payment.InstallmentID, "settlement")

		assert.NoError(t, err)
	
		mockPaymentQueryRepo.AssertExpectations(t)
		mockInstallmentCommandRepo.AssertExpectations(t)
	})
	

	t.Run("Payment Not Found", func(t *testing.T) {

		payment := domain.Payment{
			InstallmentID: "installment123",
		}

		mockPaymentQueryRepo.On("GetPaymentByInstallmentID", payment.InstallmentID).Return(domain.Payment{}, errors.New("payment not found"))

		err := paymentUsecase.UpdatePaymentStatus(payment.InstallmentID, "settlement")

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "payment not found")
		mockPaymentQueryRepo.AssertExpectations(t)
	})
}
