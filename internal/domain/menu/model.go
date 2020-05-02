package menu

import (
	"food-hub-api/internal/database"
	"food-hub-api/internal/domain/food"
	uuid "github.com/satori/go.uuid"
)

type Menu struct {
	database.Base
	Name string `gorm:"type:varchar(100);not_null;"`
	RestaurantID uuid.UUID  `gorm:"type:uuid;not_null;"`
	UserID uuid.UUID `gorm:"type:uuid;not_null;"`
	Foods []food.Food
}
