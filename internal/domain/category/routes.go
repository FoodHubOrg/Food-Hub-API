package restaurant

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/restaurant/{categoryID}",
		negroni.New(negroni.HandlerFunc(handler.Create))).Methods("POST")
	router.Handle("/restaurant/{restaurantID}",
		negroni.New(negroni.HandlerFunc(handler.Update))).Methods("PUT")
	router.Handle("/restaurant",
		negroni.New(negroni.HandlerFunc(handler.FindAll))).Methods("GET")
	router.Handle("/restaurant/{restaurantID}",
		negroni.New(negroni.HandlerFunc(handler.FindById))).Methods("GET")
	router.Handle("/restaurant/{restaurantID}",
		negroni.New(negroni.HandlerFunc(handler.Delete))).Methods("DELETE")
	return router
}
