package database

import (
	"log"

	"gorm.io/gorm"
	eu "xyz-finance-api/internal/user/entity"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(
		&eu.User{},
	)

	migrator := db.Migrator()
	tables := []string{"users","loans"}
	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}
	log.Println("all tables were successfully migrated")
}
