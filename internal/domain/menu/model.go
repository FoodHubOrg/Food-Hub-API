package menu

import (
	"Food-Hub-API/internal/database"
	"Food-Hub-API/internal/domain/food"
	uuid "github.com/satori/go.uuid"
)

type Menu struct {
	database.Base
	RestaurantID uuid.UUID  `gorm:"type:uuid;not_null;unique_index"`
	UserID uuid.UUID `gorm:"type:uuid;not_null;"`
	Foods []food.Food
}
