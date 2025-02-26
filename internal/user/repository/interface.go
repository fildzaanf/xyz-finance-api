package repository

import (
	"mime/multipart"
	"xyz-finance-api/internal/user/domain"
)

type UserCommandRepositoryInterface interface {
	RegisterUser(user domain.User, ktpPhoto *multipart.FileHeader, selfiePhoto *multipart.FileHeader) (domain.User, error)
	LoginUser(email, password string) (domain.User, error)
}

type UserQueryRepositoryInterface interface {
	GetUserByID(id string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
}
