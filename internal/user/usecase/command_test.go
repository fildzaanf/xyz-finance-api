package usecase

import (
	"errors"
	"testing"
	"xyz-finance-api/internal/user/domain"
	"xyz-finance-api/pkg/constant"
	mocks "xyz-finance-api/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestRegisterUser(t *testing.T) {
	mockUserCommandRepo := new(mocks.UserCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)
	userCommandUsecase := NewUserCommandUsecase(mockUserCommandRepo, mockUserQueryRepo)
	t.Run("Valid Register User", func(t *testing.T) {
		user := &domain.User{
			Email:           "test@example.com",
			Password:        "validpassword",
			ConfirmPassword: "validpassword",
			FullName:        "John Doe",
			Nik:             "1234567890",
			LegalName:       "John Doe",
			BirthPlace:      "City",
			BirthDate:       "1990-01-01",
			Salary:          5000000,
			KtpPhoto:        "ktp.jpg",
			SelfiePhoto:     "selfie.jpg",
		}

		mockUserQueryRepo.On("GetUserByEmail", user.Email).Return(domain.User{}, errors.New(constant.ERROR_EMAIL_UNREGISTERED))

		mockUserCommandRepo.On("RegisterUser", mock.Anything, mock.Anything, mock.Anything).Return(*user, nil)

		registeredUser, err := userCommandUsecase.RegisterUser(*user, nil, nil)

		mockUserQueryRepo.AssertExpectations(t)
		mockUserCommandRepo.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, user.Email, registeredUser.Email)
	})

	t.Run("Register User With Empty Data", func(t *testing.T) {
		user := domain.User{
			Email:           "",
			Password:        "",
			ConfirmPassword: "",
			FullName:        "",
			LegalName:       "",
			Nik:             "",
			BirthPlace:      "",
			BirthDate:       "",
			Salary:          0,
			KtpPhoto:        "",
			SelfiePhoto:     "",
		}

		_, err := userCommandUsecase.RegisterUser(user, nil, nil)

		assert.Error(t, err)
		mockUserCommandRepo.AssertExpectations(t)
	})

	t.Run("Register User With Existing Email", func(t *testing.T) {
		user := domain.User{
			Email:           "fildzanna@example.com",
			Password:        "securepassword123",
			ConfirmPassword: "securepassword123",
			FullName:        "John Doe",
			Nik:             "1234567890",
			LegalName:       "John Doe",
			BirthPlace:      "City",
			BirthDate:       "1990-01-01",
			Salary:          5000000,
			KtpPhoto:        "ktp.jpg",
			SelfiePhoto:     "selfie.jpg",
		}

		mockUserQueryRepo.On("GetUserByEmail", user.Email).Return(domain.User{}, nil)

		_, err := userCommandUsecase.RegisterUser(user, nil, nil)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_EMAIL_EXIST)

		mockUserQueryRepo.AssertExpectations(t)
	})

	t.Run("Register User With Invalid Email Fomat", func(t *testing.T) {
		user := domain.User{
			Email:           "fildzanna",
			Password:        "securepassword123",
			ConfirmPassword: "securepassword123",
			FullName:        "John Doe",
			Nik:             "1234567890",
			LegalName:       "John Doe",
			BirthPlace:      "City",
			BirthDate:       "1990-01-01",
			Salary:          5000000,
			KtpPhoto:        "ktp.jpg",
			SelfiePhoto:     "selfie.jpg",
		}

		_, err := userCommandUsecase.RegisterUser(user, nil, nil)

		assert.Error(t, err)
		mockUserCommandRepo.AssertExpectations(t)
	})

	t.Run("Register User With Password Minimal Length", func(t *testing.T) {
		user := domain.User{
			Email:           "fildzanna@example.com",
			Password:        "secure",
			ConfirmPassword: "secure",
			FullName:        "John Doe",
			Nik:             "1234567890",
			LegalName:       "John Doe",
			BirthPlace:      "City",
			BirthDate:       "1990-01-01",
			Salary:          5000000,
			KtpPhoto:        "ktp.jpg",
			SelfiePhoto:     "selfie.jpg",
		}
		_, err := userCommandUsecase.RegisterUser(user, nil, nil)

		assert.Error(t, err)
		mockUserCommandRepo.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {

	mockUserCommandRepo := new(mocks.UserCommandRepositoryInterface)
	mockUserQueryRepo := new(mocks.UserQueryRepositoryInterface)

	userCommandUsecase := NewUserCommandUsecase(mockUserCommandRepo, mockUserQueryRepo)

	t.Run("Login User With Incorrect Password", func(t *testing.T) {
		email := "fildzanna@example.com"
		password := "invalidpassword"

		user := domain.User{
			Email:    email,
			Password: "hashedpassword",
		}

		mockUserQueryRepo.On("GetUserByEmail", email).Return(user, nil)

		_, _, err := userCommandUsecase.LoginUser(email, password)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_LOGIN)

		mockUserQueryRepo.AssertExpectations(t)
		mockUserCommandRepo.AssertExpectations(t)
	})

	t.Run("Login User With Correct Credential", func(t *testing.T) {
		email := "fildzanna@example.com"
		password := "securepassword123"

		user := domain.User{
			Email:    email,
			Password: "hashedpassword",
		}

		mockUserQueryRepo.On("GetUserByEmail", email).Return(user, nil)

		_, _, err := userCommandUsecase.LoginUser(email, password)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_LOGIN)

		mockUserQueryRepo.AssertExpectations(t)
		mockUserCommandRepo.AssertExpectations(t)
	})

	t.Run("Login User With Unregistered Email", func(t *testing.T) {
		email := "unregistered@example.com"
		password := "validpassword"


		mockUserQueryRepo.On("GetUserByEmail", email).Return(domain.User{}, errors.New(constant.ERROR_EMAIL_UNREGISTERED))

		_, _, err := userCommandUsecase.LoginUser(email, password)

		assert.NotNil(t, err)
		assert.EqualError(t, err, constant.ERROR_EMAIL_UNREGISTERED)

		mockUserQueryRepo.AssertExpectations(t)
	})

	t.Run("Login User With Empty Data", func(t *testing.T) {
		user := domain.User{
			Email:           "",
			Password:        "",
		}

		_, _, err := userCommandUsecase.LoginUser(user.Email, user.Password)

		assert.Error(t, err)
		mockUserCommandRepo.AssertExpectations(t)
	})

	t.Run("Login User With Invalid Email Fomat", func(t *testing.T) {
		user := domain.User{
			Email:           "fildzanna",
			Password:        "securepassword123",
		}

		_, _, err := userCommandUsecase.LoginUser(user.Email, user.Password)

		assert.Error(t, err)
		mockUserCommandRepo.AssertExpectations(t)
	})
	


}
