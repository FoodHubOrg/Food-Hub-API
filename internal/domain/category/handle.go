package category

import (
	"food-hub-api/internal/domain/restaurant"
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
	var category restaurant.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := s.service.Create(&category)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var category restaurant.Category
	categoryID := mux.Vars(r)["categoryID"]

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	parsedID, err := uuid.FromString(categoryID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := s.service.Update(parsedID, &category)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	_ = result.ID.String()


	helpers.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"ID": categoryID,
		"Name": result.Name,
		//"Restaurants": result.Restaurants,
		"message": "successfully updated category",
	})
	return
}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	categoryID := mux.Vars(r)["categoryID"]
	parsedID, err := uuid.FromString(categoryID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = s.service.Delete(parsedID)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, map[string]string{
		"message": "category deleted successfully",
	})
	return
}

func (s *handler) FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	result, err := s.service.FindAll()
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	categoryID := mux.Vars(r)["categoryID"]
	parsedID, err := uuid.FromString(categoryID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := s.service.FindById(parsedID)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}
