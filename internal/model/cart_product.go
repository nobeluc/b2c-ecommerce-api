package model

import (
	"gorm.io/gorm"
)

type CartProduct struct {
	gorm.Model
	Quantity  uint
	CartID    uint
	ProductID uint
	Cart      Cart
	Product   Product
}