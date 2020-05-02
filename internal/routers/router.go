package routers

import (
	"food-hub-api/internal/domain/cart"
	"food-hub-api/internal/domain/category"
	"food-hub-api/internal/domain/food"
	"food-hub-api/internal/domain/menu"
	"food-hub-api/internal/domain/order"
	"food-hub-api/internal/domain/restaurant"
	"food-hub-api/internal/domain/user"
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
