package cart

import (
	"Food-Hub-API/internal/database"
	"Food-Hub-API/internal/domain/food"
	"Food-Hub-API/internal/helpers"
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
	var cart Cart
	restaurantID := mux.Vars(r)["restaurantID"]

	ids, err := helpers.ParseIDs([]string{restaurantID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails, _ := helpers.VerifyToken(r)
	cart.UserID = userDetails.ID
	cart.RestaurantID = ids[0]

	result, err := s.service.Create(&cart)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, result)
	return
}

func (s *handler) Update(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var cart Cart
	cartID := mux.Vars(r)["cartID"]
	foodID := mux.Vars(r)["foodID"]

	//parsedCartID, err := uuid.FromString(cartID)
	ids, err := helpers.ParseIDs([]string{cartID, foodID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cart.ID = ids[0]
	cart.Foods = []food.Food{{
		Base: database.Base{
			ID: ids[1],
		},
	}}

	result, err := s.service.Update(&cart)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, result)
	return
}

func (s *handler) RemoveFood(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	var cart Cart
	cartID := mux.Vars(r)["cartID"]
	foodID := mux.Vars(r)["foodID"]

	//parsedCartID, err := uuid.FromString(cartID)
	ids, err := helpers.ParseIDs([]string{cartID, foodID})
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	cart.ID = ids[0]
	cart.Foods = []food.Food{{
		Base: database.Base{
			ID: ids[1],
		},
	}}

	err = s.service.RemoveFood(&cart)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusAccepted, err)
	return
}

func (s *handler) Delete(w http.ResponseWriter, r *http.Request, n http.HandlerFunc){
	var cart Cart
	cartID := mux.Vars(r)["cartID"]

	parsedCartID, err := uuid.FromString(cartID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	//userDetails, _ := helpers.VerifyToken(r)
	cart.ID = parsedCartID
	//cart.UserID = userDetails.ID

	if err = s.service.Delete(&cart); err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "cart deleted successfully",
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
	var cart Cart
	cartID := mux.Vars(r)["cartID"]
	parsedID, err := uuid.FromString(cartID)
	if err != nil{
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	cart.ID = parsedID
	result, err := s.service.FindByID(&cart)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.JSONResponse(w, http.StatusOK, result)
	return
}

