package order

import "Food-Hub-API/internal/database"

type Order struct {
	database.Base
	Status string `gorm:"type:varchar(100);default:pending"`
	MenuID string
	FoodID string
	RestaurantID string
}
