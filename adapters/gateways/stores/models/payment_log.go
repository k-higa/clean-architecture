package models

import "database/sql"

type PaymentLog struct {
	ID          int `gorm:"primaryKey"`
	Host        sql.NullString
	ContentType sql.NullString
	Request     sql.NullString
	Response    sql.NullString
}
