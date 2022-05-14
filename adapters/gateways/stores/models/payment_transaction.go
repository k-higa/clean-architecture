package models

type PaymentTransaction struct {
	ID     int `gorm:"primaryKey"`
	Amount int
	Status string
}
