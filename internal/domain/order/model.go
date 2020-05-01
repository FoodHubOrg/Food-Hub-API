package order

import (
	"Food-Hub-API/internal/database"
	"Food-Hub-API/internal/domain/food"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	database.Base
	Street string `gorm:"type:varchar(100);not_null"`
	Number string `gorm:"type:varchar(100);not_null"`
	City string `gorm:"type:varchar(100);not_null"`
	District string `gorm:"type:varchar(100);not_null"`
	Country string `gorm:"type:varchar(100);not_null"`
	PaymentType string `gorm:"type:varchar(100);not_null"`
	Status string `gorm:"type:varchar(100);default:'pending'"`
	CartID uuid.UUID `gorm:"type:uuid;not_null;unique"`
	UserID uuid.UUID `gorm:"type:uuid;not_null"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not_null"`
	Foods []food.Food `gorm:"many2many:food_orders;"`
}
