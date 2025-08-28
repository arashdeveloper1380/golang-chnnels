package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
	SellerID    int
	Seller      *User
	Categories  []Category `gorm:"many2many:product_categories"`
}
