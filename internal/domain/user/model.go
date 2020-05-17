package user

import (
	"foodhub-api/internal/database"
	"foodhub-api/internal/domain/order"
	"foodhub-api/internal/domain/restaurant"
)

type User struct {
	database.Base
	Cpf               string `gorm:"type:varchar(100);"`
	Name              string `gorm:"type:varchar(100);not_null"`
	Email             string `gorm:"type:varchar(100);unique_index;not_null"`
	Password          string `gorm:"type:varchar(250);not_null"`
	IsDelivery        bool   `gorm:"default:false;not_null"`
	IsAdmin           bool   `gorm:"default:false;not_null"`
	IsVerified        bool 	`gorm:"default:false;not_null"`
	IsSuperAdmin 	  bool 	`gorm:"default:false;not_null"`
	IsRestaurantOwner bool `gorm:"default:false;not_null"`
	Orders            []order.Order
	Restaurants       []restaurant.Restaurant
	IsArchived  bool  `gorm:"default:false;not_null"`
}
