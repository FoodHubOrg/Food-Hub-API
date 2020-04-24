package order

import (
	"Food-Hub-API/internal/database"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	database.Base
	Status string `gorm:"type:varchar(100);default:'pending'"`
	UserID uuid.UUID `gorm:"type:uuid;not_null"`
	FoodID uuid.UUID `gorm:"type:uuid;not_null"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not_null"`
	Street string `gorm:"type:varchar(100);not_null"`
	Number string `gorm:"type:varchar(100);not_null"`
	City string `gorm:"type:varchar(100);not_null"`
	District string `gorm:"type:varchar(100);not_null"`
	Country string `gorm:"type:varchar(100);not_null"`
	PaymentType string `gorm:"type:varchar(100);not_null"`
}
