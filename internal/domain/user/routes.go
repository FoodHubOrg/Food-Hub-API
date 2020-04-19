package user

import (
	"Food-Hub-API/internal/middlewares"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/signup",
		negroni.New(negroni.HandlerFunc(handler.CreateAccount))).Methods("POST")
	router.Handle("/login",
		negroni.New(negroni.HandlerFunc(handler.Login))).Methods("POST")
	router.Handle("/create-restaurant-owner/{userID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.CreateRestaurantOwner))).Methods("POST")
	return router
}
