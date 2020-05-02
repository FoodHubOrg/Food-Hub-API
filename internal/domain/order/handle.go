package order

import (
	"Food-Hub-API/internal/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler interface {
	Checkout(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Receive(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Accept(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Decline(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindByUser(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func (s *handler) Checkout(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var order Order

	restaurantID := mux.Vars(r)["restaurantID"]
	cartID := mux.Vars(r)["cartID"]

	ids, err := helpers.ParseIDs([]string{restaurantID, cartID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	order.UserID = userDetails.ID
	order.CartID = ids[1]
	order.RestaurantID = ids[0]

	result, err := s.service.Create(&order)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Receive(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var order Order
	orderID := mux.Vars(r)["orderID"]

	ids, err := helpers.ParseIDs([]string{orderID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	order.UserID = userDetails.ID
	order.ID = ids[0]

	result, err := s.service.Update(&order, "received")
	if err != nil{
		if err.Error() == "is not owner" {
			helpers.ErrorResponse(w, http.StatusForbidden,
				"failed to perform action, please contact administration for help")
			return
		}
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, result)
	return
}

func (s *handler) Accept(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var order Order
	orderID := mux.Vars(r)["orderID"]

	ids, err := helpers.ParseIDs([]string{orderID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	order.UserID = userDetails.ID
	order.ID = ids[0]

	result, err := s.service.Update(&order, "accepted")
	if err != nil{
		if err.Error() == "is not owner" {
			helpers.ErrorResponse(w, http.StatusForbidden,
				"failed to perform action, please contact administration for help")
			return
		}
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, result)
	return
}

func (s *handler) Decline(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var order Order
	orderID := mux.Vars(r)["orderID"]

	ids, err := helpers.ParseIDs([]string{orderID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	order.UserID = userDetails.ID
	order.ID = ids[0]

	result, err := s.service.Update(&order, "declined")
	if err != nil{
		if err.Error() == "is not owner" {
			helpers.ErrorResponse(w, http.StatusForbidden,
				"failed to perform action, please contact administration for help")
			return
		}
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, result)
	return
}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var order Order
	orderID := mux.Vars(r)["orderID"]

	ids, err := helpers.ParseIDs([]string{orderID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	order.ID = ids[0]
	order.UserID = userDetails.ID

	if err = s.service.Delete(&order); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "order deleted successfully",
	})
	return
}

func (s *handler) FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	result, err := s.service.FindAll()
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

func (s *handler) FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var order Order

	orderID := mux.Vars(r)["orderID"]
	ids, err := helpers.ParseIDs([]string{orderID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	order.ID = ids[0]
	order.UserID = userDetails.ID

	result, err := s.service.FindById(&order, "order")
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}


func (s *handler) FindByUser(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var order Order

	userID := mux.Vars(r)["userID"]
	ids, err := helpers.ParseIDs([]string{userID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//userDetails, _ := helpers.VerifyToken(r)
	order.UserID = ids[0]

	result, err := s.service.FindById(&order, "user")
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}