package user

import (
	"Food-Hub-API/internal/helpers"
	"encoding/json"
	"net/http"
)

// Methods to be consumed by handler
type UserHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	Login(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type CreatedUser struct {
	Name string
	Email string
	Message string
	Token string
}

// Implements userUseCaseMethods -> Interface
// This will enable us send data to the UseCaseLayer
// Implement service layer
// Struct implements UserHandler
type userHandler struct {
	service UserService
}

// Implement UserHandler Interface
// Returns a Handler interface
func NewUserHandler(service UserService) UserHandler {
	return &userHandler{
		service,
	}
}

func (u *userHandler) CreateAccount(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

	token, err := helpers.CreateToken(user)
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

func (u *userHandler) Login(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
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

	token, err := helpers.CreateToken(user)
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

