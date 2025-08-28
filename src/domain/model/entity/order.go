package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	User   *User
	Items  []OrderItem
	Total  float64
	Status string // pending, paid
}
