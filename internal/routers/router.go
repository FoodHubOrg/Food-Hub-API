package routers

import (
	"Food-Hub-API/internal/domain/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func InitRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	router = user.SetUserRouters(router, db)
	return router
}
