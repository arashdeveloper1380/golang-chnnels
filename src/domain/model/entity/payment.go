package entity

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID       *uint
	Order         *Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Method        string
	Amount        float64
	TransactionID string
	Status        string
}
