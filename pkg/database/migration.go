package database

import (
	"log"

	"gorm.io/gorm"
	eu "xyz-finance-api/internal/user/entity"
	el "xyz-finance-api/internal/loan/entity"
	et "xyz-finance-api/internal/transaction/entity"
	ei "xyz-finance-api/internal/installment/entity"
	ep "xyz-finance-api/internal/payment/entity"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&eu.User{},
		&el.Loan{},
		&et.Transaction{},
		&ei.Installment{},
		&ep.Payment{},

	)

	migrator := db.Migrator()
	tables := []string{"users","loans","transactions","installments", "payments"}
	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}
	log.Println("all tables were successfully migrated")
}
