package usecase

import (
	"mime/multipart"
	"xyz-finance-api/internal/user/domain"
)

type UserCommandUsecaseInterface interface {
	RegisterUser(user domain.User, ktpPhoto *multipart.FileHeader, selfiePhoto *multipart.FileHeader) (domain.User, error)
	LoginUser(email, password string) (domain.User, string, error)
}

type UserQueryUsecaseInterface interface {
	GetUserByID(id string) (domain.User, error)
}
