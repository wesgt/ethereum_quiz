package utils

import (
	"database/sql"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB(t *testing.T, sqlDB *sql.DB) *gorm.DB {
	// create dialector
	dialector := mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      sqlDB,
		DriverName:                "mysql",
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm connection", err)
	}

	return db
}
