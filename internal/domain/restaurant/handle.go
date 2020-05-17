package restaurant

import (
	"foodhub-api/internal/database"
	"foodhub-api/internal/helpers"
	"foodhub-api/internal/middlewares/validations"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"reflect"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindAll(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	FindById(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
	RemoveCategory(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
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
	var restaurant Restaurant

	result := context.Get(r, "restaurant")
	rest := reflect.ValueOf(result)
	restaurant.Cover = rest.FieldByName("Cover").String()
	restaurant.Name = rest.FieldByName("Name").String()
	restaurant.Location = rest.FieldByName("Location").String()
	categories := rest.FieldByName("Categories").Interface().([]validations.Category)

	for _, v := range categories {
		restaurant.Categories = append(restaurant.Categories, Category{Name:v.Name})
	}

	userDetails, _ := helpers.VerifyToken(r)
	restaurant.UserID = userDetails.ID

	result, err := s.service.Create(&restaurant)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var restaurant Restaurant
	restaurantID := mux.Vars(r)["restaurantID"]

	ids, err := helpers.ParseIDs([]string{restaurantID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	result := context.Get(r, "restaurant")
	rest := reflect.ValueOf(result)
	restaurant.Cover = rest.FieldByName("Cover").String()
	restaurant.Name = rest.FieldByName("Name").String()
	restaurant.Time = rest.FieldByName("Time").String()
	restaurant.Location = rest.FieldByName("Location").String()
	categories := rest.FieldByName("Categories").Interface().([]validations.Category)

	for _, v := range categories {
		restaurant.Categories = append(restaurant.Categories, Category{Name:v.Name})
	}

	userDetails, _ := helpers.VerifyToken(r)
	restaurant.UserID = userDetails.ID
	restaurant.ID = ids[0]

	entity, err := s.service.Update(&restaurant)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, entity)
	return
}

func (s *handler) RemoveCategory(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var restaurant Restaurant
	restaurantID := mux.Vars(r)["restaurantID"]
	categoryID := mux.Vars(r)["categoryID"]

	ids, err := helpers.ParseIDs([]string{restaurantID, categoryID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	category := Category{
		Base:        database.Base{
			ID:ids[1],
		},
	}

	userDetails, _ := helpers.VerifyToken(r)
	restaurant.UserID = userDetails.ID
	restaurant.ID = ids[0]
	restaurant.Categories = []Category{category}

	_, err = s.service.RemoveCategory(&restaurant)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, map[string]string{
		"Message": "successfully removed tag",
	})
	return
}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var restaurant Restaurant
	restaurantID := mux.Vars(r)["restaurantID"]

	ids, err := helpers.ParseIDs([]string{restaurantID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	restaurant.ID = ids[0]
	restaurant.UserID = userDetails.ID

	if err = s.service.Delete(&restaurant); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, map[string]string{
		"Message": "restaurant deleted successfully",
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
	restaurantID := mux.Vars(r)["restaurantID"]
	parsedID, err := uuid.FromString(restaurantID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := s.service.FindById(parsedID)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}
