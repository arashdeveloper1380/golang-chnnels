package entity

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint
	Order     *Order
	ProductID uint
	Product   *Product
	Quantity  int
	Price     float64
}
