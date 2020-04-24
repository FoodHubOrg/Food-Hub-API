package user

import (
	"Food-Hub-API/internal/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// Methods to be consumed by handler
type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Login(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Orders(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	//RestaurantOwner(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type CreatedUser struct {
	Name string
	Email string
	Message string
	Token string
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func (u *handler) Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := u.service.Create(&user)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]interface{}{
		"id": result.ID,
		"email": result.Email,
		"name": result.Name,
		"isAdmin": result.IsAdmin,
		"isRestaurant": result.IsRestaurantOwner,
		"isDelivery": result.IsDelivery,
	}
	token, err := helpers.CreateToken(m)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdUser := CreatedUser{
		Name:result.Name,
		Email:result.Email,
		Message:"successfully signed up",
		Token:token,
	}

	helpers.JSONResponse(w, http.StatusCreated, createdUser)
	return
}

func (u *handler) Login(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := u.service.Login(&user, user.Password)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	m := map[string]interface{}{
		"id": user.ID,
		"email": user.Email,
		"name": user.Name,
		"isAdmin": user.IsAdmin,
		"isRestaurantOwner": user.IsRestaurantOwner,
		"isDelivery": user.IsDelivery,
	}
	token, err := helpers.CreateToken(m)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdUser := CreatedUser{
		Name:user.Name,
		Email:user.Email,
		Message:"successfully logged in",
		Token:token,
	}

	helpers.JSONResponse(w, http.StatusCreated, createdUser)
	return
}

func (u *handler) Orders(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User
	claims, err := helpers.VerifyToken(r)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	user.Email = claims.Email

	orders, err := u.service.FindBy(&user, "email")
}

//func (u *handler) RestaurantOwner(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
//	var user User
//
//	userID := mux.Vars(r)["userID"]
//	ids, err := helpers.ParseIDs([]string{userID})
//	if err != nil{
//		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	user.ID = parsedUserID
//	_, err = u.service.Update(&user)
//	if err != nil {
//		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	m := map[string]interface{}{
//		"message": "successfully created restaurant owner",
//	}
//
//	helpers.JSONResponse(w, http.StatusCreated, m)
//	return
//}