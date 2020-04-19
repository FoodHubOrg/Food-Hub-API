package restaurant

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type Service interface {
	Create(restaurant *Restaurant) (*Restaurant, error)
	Update(restaurant *Restaurant) (*Restaurant, error)
	Delete(restaurant *Restaurant) error
	FindAll()([]*Restaurant, error)
	FindById(id uuid.UUID)(Restaurant, error)
	//CheckUser(restaurant *Restaurant) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) CheckUser(restaurant *Restaurant) error {
	rest, err := s.repo.FindById(restaurant.ID)
	if err != nil {
		return err
	}

	if rest.UserID != restaurant.UserID {
		return fmt.Errorf("is not owner")
	}

	return nil
}

func (s service) Create(restaurant *Restaurant) (*Restaurant, error) {
	restaurant.Menu.UserID = restaurant.UserID
	restaurant.Menu.RestaurantID  = restaurant.ID
	result, err := s.repo.Create(restaurant)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(restaurant *Restaurant) (*Restaurant, error) {
	result, err := s.repo.Update(restaurant)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(restaurant *Restaurant) error {
	if err := s.CheckUser(restaurant); err != nil {
		return err
	}
	err := s.repo.Delete(restaurant)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Restaurant, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(id uuid.UUID) (Restaurant, error) {
	result, err := s.repo.FindById(id)
	if err != nil{
		return result, err
	}
	return result, nil
}