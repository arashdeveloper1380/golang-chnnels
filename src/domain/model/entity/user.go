package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string // customer, admin, seller
	Orders   []Order
	Products []Product `gorm:"foreignKey:SellerID"`
}
