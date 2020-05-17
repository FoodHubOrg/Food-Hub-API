package user

import (
	"foodhub-api/internal/middlewares"
	"foodhub-api/internal/middlewares/validations"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/user/signup",
		negroni.New(
			negroni.HandlerFunc(validations.ReturnHandler(db).InputCreateAccount),
			negroni.HandlerFunc(handler.Create))).Methods("POST")
	router.Handle("/user/{userID}/restaurant/make",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.MakeRestaurantOwner))).Methods("PATCH")
	router.Handle("/user/{userID}/restaurant/revoke",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireAdminRights),
			negroni.HandlerFunc(handler.RevokeRestaurantOwner))).Methods("PATCH")
	router.Handle("/auth/google/login",
		negroni.New(negroni.HandlerFunc(handler.GoogleLogin))).Methods("GET")
	router.Handle("/auth/google/callback",
		negroni.New(negroni.HandlerFunc(handler.GoogleCallBack))).Methods("GET")
	router.Handle("/auth/facebook/login",
		negroni.New(negroni.HandlerFunc(handler.FacebookLogin))).Methods("GET")
	router.Handle("/auth/facebook/callback",
		negroni.New(negroni.HandlerFunc(handler.GoogleCallBack))).Methods("GET")
	router.Handle("/user/login",
		negroni.New(
			negroni.HandlerFunc(validations.ReturnHandler(db).InputLogin),
			negroni.HandlerFunc(handler.Login))).Methods("POST")
	router.Handle("/user/orders",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(handler.Orders))).Methods("GET")
	router.Handle("/user/restaurants",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(handler.Restaurants))).Methods("GET")

	return router
}
