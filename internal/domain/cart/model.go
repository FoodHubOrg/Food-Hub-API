package cart

import (
	"food-hub-api/internal/database"
	"food-hub-api/internal/domain/food"
	uuid "github.com/satori/go.uuid"
)

type Cart struct {
	database.Base
	UserID uuid.UUID `gorm:"type:uuid;not_null;"`
	RestaurantID  uuid.UUID `gorm:"type:uuid;not_null;"`
	Foods []food.Food `gorm:"many2many:cart_foods;"`
}
