package restaurant

import (
	"food-hub-api/internal/database"
	"food-hub-api/internal/domain/menu"
	"food-hub-api/internal/domain/order"
	uuid "github.com/satori/go.uuid"
)

type Restaurant struct {
	database.Base
	Name string `gorm:"type:varchar(100);not_null;unique_index"`
	Location string `gorm:"type:varchar(100);not_null"`
	UserID  uuid.UUID `gorm:"type:uuid;not_null"`
	Time string `gorm:"size:100;not_null"`
	Categories []Category `gorm:"many2many:restaurant_categories;"`
	Orders []order.Order
	Menus []menu.Menu
}

type Category struct {
	database.Base
	Name  string `gorm:"not_null;unique_index"`
	Restaurants []Restaurant `gorm:"many2many:restaurant_categories;"`
}

