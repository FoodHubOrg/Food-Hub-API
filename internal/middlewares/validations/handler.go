package validations

import (
	"encoding/json"
	"foodhub-api/internal/helpers"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler interface {
	InputCreateAccount(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	InputLogin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	InputMenu(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	InputRestaurant(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	InputMessage(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	Order(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	CheckOwner(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	CheckMenuOwner(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func(s *handler) InputCreateAccount(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := validation.ValidateStruct(&user,
		validation.Field(&user.Name,
			validation.Required.Error("name is required"),
			validation.Length(1, 50)),
		validation.Field(&user.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("please provide a valid email")),
		validation.Field(&user.Password, validation.Required.Error("password is required")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	context.Set(r, "user", user)

	next(w, r)

	return
}

func(s *handler) Order(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var order Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := validation.ValidateStruct(&order,
		validation.Field(&order.Street,
			validation.Required.Error("please provide a street")),
		validation.Field(&order.PaymentType,
			validation.Required.Error("please provide a payment method")),
		validation.Field(&order.City,
			validation.Required.Error("please provide a city")),
		validation.Field(&order.District,
			validation.Required.Error("please provide a district")),
		validation.Field(&order.Country, validation.Required.Error("please provide country")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	context.Set(r, "order", order)

	next(w, r)

	return
}

func (s *handler) InputLogin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required.Error("email is required")),
		validation.Field(&user.Password, validation.Required.Error("password is required")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	context.Set(r, "user", user)

	next(w, r)

	return
}

func (s *handler) InputRestaurant(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var restaurant Restaurant

	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := validation.ValidateStruct(&restaurant,
		validation.Field(&restaurant.Name,
			validation.Required.Error("please provide a name for the restaurant")),
		validation.Field(&restaurant.Location,
			validation.Required.Error("please provide this restaurants location")),
		validation.Field(&restaurant.Cover,
			validation.Required.Error("please provide restaurant cover photo")),
		validation.Field(&restaurant.Categories,
			validation.Required.Error("please provide at least one category")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	context.Set(r, "restaurant", restaurant)

	next(w, r)

	return
}


func (s *handler) CheckOwner(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	var restaurant Restaurant
	restaurantID := mux.Vars(r)["restaurantID"]

	ids, err := helpers.ParseIDs([]string{restaurantID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	auth, _ := helpers.VerifyToken(r)
	restaurant.ID = ids[0]
	restaurant.UserID = auth.ID

	if err := s.service.CheckOwner(&restaurant); err != nil {
		helpers.ErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}

	next(w, r)
}

func (s *handler) CheckMenuOwner(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	var menu Menu
	menuID := mux.Vars(r)["menuID"]

	ids, err := helpers.ParseIDs([]string{menuID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	auth, _ := helpers.VerifyToken(r)
	menu.ID = ids[0]
	menu.UserID = auth.ID

	if err := s.service.CheckMenuOwner(&menu); err != nil {
		helpers.ErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}

	next(w, r)
}

func (s *handler) InputMenu(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var menu Menu

	if err := json.NewDecoder(r.Body).Decode(&menu); err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := validation.ValidateStruct(&menu,
		validation.Field(&menu.Name, validation.Required.Error("menu name is required")),
		validation.Field(&menu.Foods, validation.Required.Error("at-least one food is required")),
	)

	if err != nil{
		helpers.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	context.Set(r, "menu", menu)

	next(w, r)

	return
}


func (s *handler) InputMessage(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	panic("implement me")
}

