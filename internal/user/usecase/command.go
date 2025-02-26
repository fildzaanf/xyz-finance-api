package usecase

import (
	"errors"
	"mime/multipart"
	"xyz-finance-api/pkg/cloud"
	"xyz-finance-api/pkg/constant"
	"xyz-finance-api/pkg/crypto"
	"xyz-finance-api/pkg/middleware"
	"xyz-finance-api/pkg/validator"
	"xyz-finance-api/internal/user/domain"
	"xyz-finance-api/internal/user/repository"
)

type userCommandUsecase struct {
	userCommandRepository repository.UserCommandRepositoryInterface
	userQueryRepository   repository.UserQueryRepositoryInterface
}

func NewUserCommandUsecase(ucr repository.UserCommandRepositoryInterface, uqr repository.UserQueryRepositoryInterface) UserCommandUsecaseInterface {
	return &userCommandUsecase{
		userCommandRepository: ucr,
		userQueryRepository:   uqr,
	}
}

func (ucs *userCommandUsecase) RegisterUser(user domain.User, ktpPhoto *multipart.FileHeader, selfiePhoto *multipart.FileHeader) (domain.User, error) {

	errEmpty := validator.IsDataEmpty([]string{"email", "password", "confirm_password", "nik", "full_name", "legal_name", "birth_place", "birh_date", "salary", "ktp_photo", "selfie_photo"}, user.Email, user.Password, user.ConfirmPassword, user.Nik, user.FullName, user.LegalName, user.BirthPlace, user.BirthDate, user.Salary, user.KtpPhoto, user.SelfiePhoto)
	if errEmpty != nil {
		return domain.User{}, errEmpty
	}

	if ktpPhoto != nil {
		ktpURL, errUpload := cloud.UploadImageToS3(ktpPhoto)
		if errUpload != nil {
			return domain.User{}, errUpload
		}
		user.KtpPhoto = ktpURL
	}

	if selfiePhoto != nil {
		selfieURL, errUpload := cloud.UploadImageToS3(selfiePhoto)
		if errUpload != nil {
			return domain.User{}, errUpload
		}
		user.SelfiePhoto = selfieURL
	}

	errEmailValid := validator.IsEmailValid(user.Email)
	if errEmailValid != nil {
		return domain.User{}, errEmailValid
	}

	errLength := validator.IsMinLengthValid(10, map[string]string{"password": user.Password})
	if errLength != nil {
		return domain.User{}, errLength
	}

	_, errGetEmail := ucs.userQueryRepository.GetUserByEmail(user.Email)
	if errGetEmail == nil {
		return domain.User{}, errors.New(constant.ERROR_EMAIL_EXIST)
	}

	if user.Password != user.ConfirmPassword {
		return domain.User{}, errors.New(constant.ERROR_PASSWORD_CONFIRM)
	}

	hashedPassword, errHash := crypto.HashPassword(user.Password)
	if errHash != nil {
		return domain.User{}, errors.New(constant.ERROR_PASSWORD_HASH)
	}

	user.Password = hashedPassword

	userEntity, errRegister := ucs.userCommandRepository.RegisterUser(user, ktpPhoto, selfiePhoto)
	if errRegister != nil {
		return domain.User{}, errRegister
	}

	return userEntity, nil
}

func (ucs *userCommandUsecase) LoginUser(email, password string) (domain.User, string, error) {

	errEmpty := validator.IsDataEmpty([]string{"email", "password"}, email, password)
	if errEmpty != nil {
		return domain.User{}, "", errEmpty
	}

	errEmailValid := validator.IsEmailValid(email)
	if errEmailValid != nil {
		return domain.User{}, "", errEmailValid
	}

	userDomain, errGetEmail := ucs.userQueryRepository.GetUserByEmail(email)
	if errGetEmail != nil {
		return domain.User{}, "", errors.New(constant.ERROR_EMAIL_UNREGISTERED)
	}

	comparePassword := crypto.ComparePassword(userDomain.Password, password)
	if comparePassword != nil {
		return domain.User{}, "", errors.New(constant.ERROR_LOGIN)
	}

	token, errCreate := middleware.GenerateToken(userDomain.ID, userDomain.Role)
	if errCreate != nil {
		return domain.User{}, "", errors.New(constant.ERROR_TOKEN_GENERATE)
	}

	return userDomain, token, nil
}
