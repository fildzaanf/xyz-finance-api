package database

import (
	"fmt"
	"log"
	"xyz-finance-api/pkg/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgreSQL() *gorm.DB {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load PostgreSQL configuration: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.POSTGRESQL.POSTGRESQL_HOST,
		config.POSTGRESQL.POSTGRESQL_USER,
		config.POSTGRESQL.POSTGRESQL_PASS,
		config.POSTGRESQL.POSTGRESQL_NAME,
		config.POSTGRESQL.POSTGRESQL_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect PostgreSQL: %v", err)
	}

	Migration(db)

	logrus.Info("connected to PostgreSQL")

	return db
}
