package user

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func SetUserRouters(router *mux.Router, db *gorm.DB) *mux.Router {
	userRepo := NewUserRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)
	router.Handle("/signup",
		negroni.New(
			negroni.HandlerFunc(userHandler.CreateAccount),
		)).Methods("POST")
	router.Handle("/login",
		negroni.New(
			negroni.HandlerFunc(userHandler.Login),
		)).Methods("POST")
	return router
}
