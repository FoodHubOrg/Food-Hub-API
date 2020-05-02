package cart

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(cart *Cart) (*Cart, error)
	Update(cart *Cart) (*Cart, error)
	Delete(cart *Cart) error
	FindAll()([]*Cart, error)
	FindByID(cart *Cart)(*Cart, error)
	RemoveFood(cart *Cart) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(cart *Cart) (*Cart, error) {
	result, err := s.repo.Create(cart)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(cart *Cart) (*Cart, error) {
	result, err := s.repo.Update(cart)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) RemoveFood(cart *Cart) error {
	err := s.repo.Remove(cart)
	if err != nil{
		return err
	}
	return nil
}


func (s service) Delete(cart *Cart) error {
	err := s.repo.Delete(cart)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Cart, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindByID(cart *Cart) (*Cart, error) {
	result, err := s.repo.FindByID(cart)
	if err != nil{
		return result, err
	}
	return result, nil
}
