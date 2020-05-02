package user

import (
	"food-hub-api/internal/middlewares"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/user/signup",
		negroni.New(negroni.HandlerFunc(handler.Create))).Methods("POST")
	router.Handle("/user/login",
		negroni.New(negroni.HandlerFunc(handler.Login))).Methods("POST")
	router.Handle("/user/orders",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(handler.Orders))).Methods("GET")
	router.Handle("/user/restaurants",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(handler.Restaurants))).Methods("GET")

	return router
}
