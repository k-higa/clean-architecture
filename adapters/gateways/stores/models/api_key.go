package models

import "database/sql"

type APIKey struct {
	ID     int `gorm:"primaryKey"`
	KEY    sql.NullString
	Expire sql.NullTime
}
