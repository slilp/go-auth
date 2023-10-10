package postgres

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB() (*gorm.DB, error) {
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
		viper.GetString("db.host"), viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.database"), viper.GetString("db.port"), viper.GetString("db.mode"))

	return dsn
}
