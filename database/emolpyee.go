package database

import "database/sql"

type Emoloyee struct {
	ID   int `gorm:"primaryKey"`
	Name sql.NullString
	Age  sql.NullInt32
}
