package database

import "gorm.io/gorm"

var (
	DB *gorm.DB
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
