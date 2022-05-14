package app

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Logger echo.Logger
)

type DBManager struct {
}

func (d *DBManager) Open() {
	db, err := gorm.Open(nil)
	if err != nil {
		panic(err)
	}
	DB = db
}
