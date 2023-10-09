package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initialize() (*gorm.DB, error) {
	dsn := generateDSN()

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})

	return db, err
}

func generateDSN() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=%s",
		"localhost", "postgres", "postgres", "goauth", "5432", "disable")

	return dsn
}
