package restaurant

import (
	"foodhub-api/internal/database"
	"foodhub-api/internal/domain/menu"
	"foodhub-api/internal/domain/order"
	uuid "github.com/satori/go.uuid"
)

type Restaurant struct {
	database.Base
	Cover string `gorm:"type:text"`
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

