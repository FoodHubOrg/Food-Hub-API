package restaurant

import (
	"Food-Hub-API/internal/database"
	"Food-Hub-API/internal/domain/category"
	"Food-Hub-API/internal/domain/menu"
	uuid "github.com/satori/go.uuid"
)

type Restaurant struct {
	database.Base
	Name string `gorm:"type:varchar(100);not_null;unique_index"`
	Location string `gorm:"type:varchar(100);not_null"`
	UserID   uuid.UUID `gorm:"type:uuid;not_null"`
	Menu     menu.Menu
	Categories []category.Category `gorm:"many2many:category_restaurants;"`
}

