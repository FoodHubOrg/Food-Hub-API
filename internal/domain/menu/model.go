package menu

import (
	"foodhub-api/internal/database"
	"foodhub-api/internal/domain/food"
	uuid "github.com/satori/go.uuid"
)

type Menu struct {
	database.Base
	Name string `gorm:"type:varchar(100);not_null;"`
	RestaurantID uuid.UUID  `gorm:"type:uuid;not_null;"`
	UserID uuid.UUID `gorm:"type:uuid;not_null;"`
	Foods []food.Food  `gorm:"many2many:menu_foods;"`
}
