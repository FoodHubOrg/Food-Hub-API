package restaurant

import (
	uuid "github.com/satori/go.uuid"
)

type Service interface {
	Create(restaurant *Restaurant) (*Restaurant, error)
	Update(restaurant *Restaurant) (*Restaurant, error)
	Delete(restaurant *Restaurant) error
	FindAll()([]*Restaurant, error)
	FindById(id uuid.UUID)(Restaurant, error)
	RemoveCategory(restaurant *Restaurant) (*Restaurant, error)
}

type service struct {
	repo Repository
}


func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(restaurant *Restaurant) (*Restaurant, error) {
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

func (s service) RemoveCategory(restaurant *Restaurant) (*Restaurant, error) {
	result, err := s.repo.RemoveCategory(restaurant)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(restaurant *Restaurant) error {
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