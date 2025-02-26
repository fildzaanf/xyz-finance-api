package router

import (
	"xyz-finance-api/internal/user/handler"
	"xyz-finance-api/internal/user/repository"
	"xyz-finance-api/internal/user/usecase"
	"xyz-finance-api/pkg/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoutes(user *echo.Group, db *gorm.DB) {
	userQueryRepository := repository.NewUserQueryRepository(db)
	userCommandRepository := repository.NewUserCommandRepository(db)

	userQueryUsecase := usecase.NewUserQueryUsecase(userCommandRepository, userQueryRepository)
	userCommandUsecase := usecase.NewUserCommandUsecase(userCommandRepository, userQueryRepository)

	userHandler := handler.NewUserHandler(userCommandUsecase, userQueryUsecase)

	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.LoginUser)
	user.GET("/:user_id", userHandler.GetUserByID, middleware.JWTMiddleware(false))
}

