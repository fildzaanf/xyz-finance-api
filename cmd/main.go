package main

import (
	lr "xyz-finance-api/internal/loan/router"
	tr "xyz-finance-api/internal/transaction/router"
	ur "xyz-finance-api/internal/user/router"
	ir "xyz-finance-api/internal/installment/router"
	pr "xyz-finance-api/internal/payment/router"
	"xyz-finance-api/pkg/config"
	"xyz-finance-api/pkg/database"
	"xyz-finance-api/pkg/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	user := e.Group("/users")
	ur.UserRoutes(user, db)

	loan := e.Group("/loans")
	lr.LoanRoutes(loan, db)

	transaction := e.Group("/transactions")
	tr.TransactionRoutes(transaction, db)

	installment := e.Group("/installments")
	ir.InstallmentRoutes(installment, db)

	payment := e.Group("/payments")
	pr.PaymentRoutes(payment, db)

}

func main() {
	godotenv.Load()
	config, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load configuration: %v", err)
	}

	//pdb := database.ConnectPostgreSQL()
	mdb := database.ConnectMySQL()

	e := echo.New()

	middleware.RemoveTrailingSlash(e)
	middleware.Logger(e)
	middleware.RateLimiter(e)
	middleware.Recover(e)
	middleware.CORS(e)

	SetupRoutes(e, mdb)

	host := config.SERVER.SERVER_HOST
	port := config.SERVER.SERVER_PORT
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "8000"
	}
	address := host + ":" + port

	logrus.Info("server is running on address ", address)
	if err := e.Start(address); err != nil {
		logrus.Fatalf("error starting server: %v", err)
	}
}
