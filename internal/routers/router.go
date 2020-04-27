package routers

import (
	"Food-Hub-API/internal/domain/cart"
	"Food-Hub-API/internal/domain/category"
	"Food-Hub-API/internal/domain/food"
	"Food-Hub-API/internal/domain/menu"
	"Food-Hub-API/internal/domain/order"
	"Food-Hub-API/internal/domain/restaurant"
	"Food-Hub-API/internal/domain/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func InitRoutes(db *gorm.DB) *mux.Router {
	// Migrations
	db.AutoMigrate(
		&order.Order{},
		&user.User{},
		&restaurant.Restaurant{},
		&menu.Menu{},
		&food.Food{},
		&restaurant.Category{},
		&cart.Cart{},
		)

	router := mux.NewRouter()
	newRouter := router.PathPrefix("/api").Subrouter()
	newRouter = category.Routes(newRouter, db)
	newRouter = restaurant.Routes(newRouter, db)
	newRouter = menu.Routes(newRouter, db)
	newRouter = order.Routes(newRouter, db)
	newRouter = food.Routes(newRouter, db)
	newRouter = user.Routes(newRouter, db)
	return newRouter
}
