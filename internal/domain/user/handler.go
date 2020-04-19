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
	CreateAccount(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Login(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	CreateRestaurantOwner(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
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

func (u *handler) CreateAccount(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err := u.service.CreateAccount(&user)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]interface{}{
		"id": user.ID,
		"email": user.Email,
		"name": user.Name,
		"isAdmin": user.IsAdmin,
		"isRestaurant": user.IsRestaurantOwner,
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

func (u *handler) CreateRestaurantOwner(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User

	userIDStr := mux.Vars(r)["userID"]
	parsedUserID, err := uuid.FromString(userIDStr)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.ID = parsedUserID
	err = u.service.Update(&user)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]interface{}{
		"message": "successfully created restaurant owner",
	}

	helpers.JSONResponse(w, http.StatusCreated, m)
	return
}