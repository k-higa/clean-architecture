package app

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB     *gorm.DB
	Logger echo.Logger
)

type DBManager struct {
}

func (d *DBManager) Open() {
	dbDsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dbDsn}))
	if err != nil {
		panic(err)
	}
	DB = db
}
