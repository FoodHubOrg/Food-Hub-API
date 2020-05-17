package routers

import (
	"foodhub-api/internal/domain/cart"
	"foodhub-api/internal/domain/category"
	"foodhub-api/internal/domain/food"
	"foodhub-api/internal/domain/menu"
	"foodhub-api/internal/domain/order"
	"foodhub-api/internal/domain/restaurant"
	"foodhub-api/internal/domain/user"
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
	newRouter = cart.Routes(newRouter, db)
	return newRouter
}
