package database

import (
	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate()

	migrator := db.Migrator()
	tables := []string{"users", "loans", "transactions", "payments"}
	for _, table := range tables {
		if !migrator.HasTable(table) {
			log.Fatalf("table %s was not successfully created", table)
		}
	}
	log.Println("all tables were successfully migrated")
}
