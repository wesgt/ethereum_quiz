package utils

import (
	"fmt"
	"time"

	"example.com/portto/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type DBPool struct {
	Dataset *gorm.DB
}

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func InitDB(cfg *config.DatabaseConfig) (err error) {
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(300)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}
