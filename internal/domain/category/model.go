package restaurant

import (
	"Food-Hub-API/internal/database"
	"Food-Hub-API/internal/domain/restaurant"
)

type Category struct {
	database.Base
	Name string
	Restaurant []restaurant.Restaurant
}
