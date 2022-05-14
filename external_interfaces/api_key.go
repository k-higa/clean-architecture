package external_interfaces

import "database/sql"

type APIKey struct {
	ID     int `gorm:"primaryKey"`
	KEY    sql.NullString
	Expire sql.NullTime
}
