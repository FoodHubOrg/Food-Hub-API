package menu

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
	router.Handle("/menu/{restaurantID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).CheckOwner),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputMenu),
			negroni.HandlerFunc(handler.Create))).Methods("POST")
	router.Handle("/menu/{menuID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).CheckMenuOwner),
			negroni.HandlerFunc(validations.ReturnHandler(db).InputMenu),
			negroni.HandlerFunc(handler.Update))).Methods("PUT")
	router.Handle("/menu/{menuID}/food/{foodID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).CheckMenuOwner),
			negroni.HandlerFunc(handler.RemoveFood))).Methods("PATCH")
	router.Handle("/menu",
		negroni.New(negroni.HandlerFunc(handler.FindAll))).Methods("GET")
	router.Handle("/menu/{menuID}",
		negroni.New(negroni.HandlerFunc(handler.FindById))).Methods("GET")
	router.Handle("/menu/{menuID}",
		negroni.New(
			negroni.HandlerFunc(middlewares.RequireTokenAuthentication),
			negroni.HandlerFunc(middlewares.RequireOwnerRights),
			negroni.HandlerFunc(validations.ReturnHandler(db).CheckMenuOwner),
			negroni.HandlerFunc(handler.Delete))).Methods("DELETE")
	return router
}
