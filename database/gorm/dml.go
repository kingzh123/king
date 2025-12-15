package gorm

import (
	"king/gorm/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPostgresDB() (db *gorm.DB, err error) {
	dsn := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable"
	loggers, err := model.NewDateFileLogger("../logs", "sql", logger.Info)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggers,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
