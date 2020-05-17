package menu

import (
	"foodhub-api/internal/database"
	"foodhub-api/internal/domain/food"
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
	RemoveFood(w http.ResponseWriter, r *http.Request, n http.HandlerFunc)
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
	var menu Menu
	restaurantID := mux.Vars(r)["restaurantID"]

	ids, err := helpers.ParseIDs([]string{restaurantID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	result := context.Get(r, "menu")
	men := reflect.ValueOf(result)
	menu.Name = men.FieldByName("Name").String()
	foods := men.FieldByName("Foods").Interface().([]validations.Food)

	for _, v := range foods {
		menu.Foods = append(menu.Foods, food.Food{Name:v.Name, Price:v.Price})
	}

	userDetails, _ := helpers.VerifyToken(r)
	menu.UserID = userDetails.ID
	menu.RestaurantID = ids[0]

	entity, err := s.service.Create(&menu)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, entity)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var menu Menu
	menuID := mux.Vars(r)["menuID"]

	ids, err := helpers.ParseIDs([]string{menuID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	result := context.Get(r, "menu")
	men := reflect.ValueOf(result)
	menu.Name = men.FieldByName("Name").String()
	foods := men.FieldByName("Foods").Interface().([]validations.Food)

	for _, v := range foods {
		menu.Foods = append(menu.Foods, food.Food{Name:v.Name, Price:v.Price})
	}

	userDetails, _ := helpers.VerifyToken(r)
	menu.UserID = userDetails.ID
	menu.ID = ids[0]

	entity, err := s.service.Update(&menu)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, entity)
	return
}

func (s *handler) RemoveFood(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var menu Menu
	menuID := mux.Vars(r)["menuID"]
	foodID := mux.Vars(r)["foodID"]

	ids, err := helpers.ParseIDs([]string{menuID, foodID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	foody := food.Food{
		Base:        database.Base{
			ID:ids[1],
		},
	}

	userDetails, _ := helpers.VerifyToken(r)
	menu.UserID = userDetails.ID
	menu.ID = ids[0]
	menu.Foods = []food.Food{foody}

	err = s.service.RemoveFood(&menu)
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
	var menu Menu
	menuID := mux.Vars(r)["menuID"]

	ids, err := helpers.ParseIDs([]string{menuID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	menu.ID = ids[0]
	menu.UserID = userDetails.ID

	if err = s.service.Delete(&menu); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "menu deleted successfully",
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
	var menu Menu
	menuID := mux.Vars(r)["menuID"]
	parsedID, err := uuid.FromString(menuID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	menu.ID = parsedID
	result, err := s.service.FindById(&menu)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

