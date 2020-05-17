package user

import (
	"encoding/json"
	"fmt"
	"foodhub-api/internal/helpers"
	"github.com/dchest/uniuri"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
	"reflect"
)

// Methods to be consumed by handler
type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Login(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Orders(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Restaurants(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	GoogleLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	GoogleCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FacebookLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FacebookCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	MakeRestaurantOwner(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	RevokeRestaurantOwner(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type CreatedUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
	Token string `json:"token"`
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
}

var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:5500/api/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

var faceBookOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:5500/api/auth/facebook/callback",
	ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
	ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
	Scopes:       []string{"email"},
	Endpoint:     facebook.Endpoint,
}


func (u *handler) FacebookLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	oauthStateString := uniuri.New()
	url := faceBookOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (u *handler) FacebookCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user GoogleUser
	content, err := helpers.GetUserDataFromMedia(r, os.Getenv("FACEBOOK_URL"), faceBookOAuthConfig)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, "failed to get user from token")
		return
	}

	err = json.Unmarshal(content, &user)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "failed to get user details")
		return
	}

	result, err := u.service.Create(&User{
		Name:user.Name,
		IsVerified:user.VerifiedEmail,
		Email:user.Email,
		Password: uniuri.New(),
	}, "social")
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	m := map[string]interface{}{
		"ID": result.ID,
		"Email": result.Email,
		"Name": result.Name,
	}

	token, err := helpers.CreateToken(m)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	url := fmt.Sprintf("%s?token=%s", os.Getenv("FRONTEND_URL"), token)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (u *handler) GoogleLogin(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	oauthStateString := uniuri.New()
	url := googleOAuthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (u *handler) GoogleCallBack(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user GoogleUser
	content, err := helpers.GetUserDataFromMedia(r, os.Getenv("GOOGLE_URL"), googleOAuthConfig)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, "failed to get user from token")
		return
	}

	err = json.Unmarshal(content, &user)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "failed to get user details")
		return
	}

	result, err := u.service.Create(&User{
		Name:user.Name,
		IsVerified:user.VerifiedEmail,
		Email:user.Email,
		Password: uniuri.New(),
	}, "social")

	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	m := map[string]interface{}{
		"ID": result.ID,
		"Email": result.Email,
		"Name": result.Name,
	}

	token, err := helpers.CreateToken(m)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, "failed to authenticate user")
		return
	}

	url := fmt.Sprintf("%s?token=%s", os.Getenv("FRONTEND_URL"), token)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return
}

func (u *handler) Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User
	result := context.Get(r, "user")
	usr := reflect.ValueOf(result)

	user.Email = usr.FieldByName("Email").String()
	user.Name = usr.FieldByName("Name").String()
	user.Password = usr.FieldByName("Password").String()

	entity, err := u.service.Create(&user, "system")
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]interface{}{
		"ID": entity.ID,
		"Email": entity.Email,
		"Name": entity.Name,
		"IsAdmin": entity.IsAdmin,
		"IsRestaurant": entity.IsRestaurantOwner,
		"IsDelivery": entity.IsDelivery,
	}
	token, err := helpers.CreateToken(m)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdUser := CreatedUser{
		Name:entity.Name,
		Email:entity.Email,
		Message:"successfully signed up",
		Token:token,
	}

	helpers.JSONResponse(w, http.StatusCreated, createdUser)
	return
}

func (u *handler) Login(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var user User

	result := context.Get(r, "user")
	usr := reflect.ValueOf(result)

	user.Email = usr.FieldByName("Email").String()
	user.Password = usr.FieldByName("Password").String()

	err := u.service.Login(&user, user.Password)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]interface{}{
		"ID": user.ID,
		"Email": user.Email,
		"Name": user.Name,
		"IsAdmin": user.IsAdmin,
		"IsRestaurantOwner": user.IsRestaurantOwner,
		"IsDelivery": user.IsDelivery,
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

	helpers.JSONResponse(w, http.StatusOK, createdUser)
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
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"Orders": orders.Orders,
	})
	return
}

func (u *handler) Restaurants(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User

	claims, err := helpers.VerifyToken(r)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.Email = claims.Email

	restaurants, err := u.service.FindBy(&user, "email")
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"Restaurants": restaurants.Restaurants,
	})
	return
}

func (u *handler) MakeRestaurantOwner(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User

	userID := mux.Vars(r)["userID"]
	ids, err := helpers.ParseIDs([]string{userID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user.ID = ids[0]
	result, err := u.service.Update(&user,"makeRestaurantOwner")
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]interface{}{
		"Message": fmt.Sprintf("successfully made %s a restaurant owner", result.Name),
	}

	helpers.JSONResponse(w, http.StatusCreated, m)
	return
}

func (u *handler) RevokeRestaurantOwner(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var user User

	userID := mux.Vars(r)["userID"]
	ids, err := helpers.ParseIDs([]string{userID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user.ID = ids[0]
	_, err = u.service.Update(&user,"makeRestaurantOwner")
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]interface{}{
		"Message": fmt.Sprintf("successfully revoked restaurant owner rights to %s", user.Name),
	}

	helpers.JSONResponse(w, http.StatusCreated, m)
	return
}
