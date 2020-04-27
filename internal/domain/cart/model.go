package cart

import (
	"Food-Hub-API/internal/database"
	"Food-Hub-API/internal/domain/order"
)

type Cart struct {
	database.Base
	Street string `gorm:"type:varchar(100);not_null"`
	Number string `gorm:"type:varchar(100);not_null"`
	City string `gorm:"type:varchar(100);not_null"`
	District string `gorm:"type:varchar(100);not_null"`
	Country string `gorm:"type:varchar(100);not_null"`
	PaymentType string `gorm:"type:varchar(100);not_null"`
	Orders []order.Order
}
