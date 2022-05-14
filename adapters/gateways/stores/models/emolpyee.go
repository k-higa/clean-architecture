package models

import "database/sql"

type Employee struct {
	ID   int `gorm:"primaryKey"`
	Name sql.NullString
	Age  sql.NullInt32
}
