package order

import (
	"Food-Hub-API/internal/database"
	"Food-Hub-API/internal/domain/food"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	database.Base
	Status string `gorm:"type:varchar(100);default:'pending'"`
	CartID uuid.UUID `gorm:"type:uuid;not_null;unique"`
	UserID uuid.UUID `gorm:"type:uuid;not_null"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not_null"`
	Food food.Food
}
