package user

import (
	"Food-Hub-API/internal/database"
	"Food-Hub-API/internal/domain/order"
	"Food-Hub-API/internal/domain/restaurant"
)

type User struct {
	database.Base
	Cpf string `gorm:"type:varchar(100);"`
	Name string `gorm:"type:varchar(100);not_null"`
	Email string `gorm:"type:varchar(100);unique_index;not_null"`
	Password string `gorm:"type:varchar(250);not_null"`
	IsRestaurantOwner bool `gorm:"default:false;not_null"`
	IsDelivery bool `gorm:"default:false;not_null"`
	IsAdmin bool `gorm:"default:false;not_null"`
	Orders []order.Order
	Restaurants []restaurant.Restaurant
}