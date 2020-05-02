package food

import (
	"food-hub-api/internal/helpers"
	"encoding/json"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func (s *handler) Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var food Food
	menuID := mux.Vars(r)["menuID"]

	parsedMenuID, err := uuid.FromString(menuID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&food); err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//userDetails, _ := helpers.VerifyToken(r)
	food.MenuID = parsedMenuID

	result, err := s.service.Create(&food)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var food Food
	foodID := mux.Vars(r)["foodID"]

	parsedFoodID, err := uuid.FromString(foodID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&food); err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//userDetails, _ := helpers.VerifyToken(r)
	//food.UserID = userDetails.ID
	food.ID = parsedFoodID

	result, err := s.service.Update(&food)
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
	var food Food
	foodID := mux.Vars(r)["foodID"]

	parsedFoodID, err := uuid.FromString(foodID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	//userDetails, _ := helpers.VerifyToken(r)
	food.ID = parsedFoodID
	//food.UserID = userDetails.ID

	if err = s.service.Delete(&food); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "food deleted successfully",
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
	var food Food
	foodID := mux.Vars(r)["foodID"]
	parsedID, err := uuid.FromString(foodID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	food.ID = parsedID
	result, err := s.service.FindById(&food)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

