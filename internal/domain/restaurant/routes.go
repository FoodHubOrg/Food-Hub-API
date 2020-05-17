package restaurant

import (
	"foodhub-api/internal/middlewares"
	"foodhub-api/internal/middlewares/validations"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
)

func Routes(router *mux.Router, db *gorm.DB) *mux.Router {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	router.Handle("/restaurant",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputRestaurant),
			negroni.HandlerFunc(handler.Create))).Methods("POST")
	router.Handle("/restaurant/{restaurantID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).CheckOwner),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputRestaurant),
			negroni.HandlerFunc(handler.Update))).Methods("PUT")
	router.Handle("/restaurant",
		negroni.New(negroni.HandlerFunc(handler.FindAll))).Methods("GET")
	router.Handle("/restaurant/{restaurantID}",
		negroni.New(negroni.HandlerFunc(handler.FindById))).Methods("GET")
	router.Handle("/restaurant/{restaurantID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).CheckOwner),
			negroni.HandlerFunc(handler.Delete))).Methods("DELETE")
	router.Handle("/restaurant/{restaurantID}/category/{categoryID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).CheckOwner),
			negroni.HandlerFunc(handler.RemoveCategory))).Methods("PATCH")
	return router
}
