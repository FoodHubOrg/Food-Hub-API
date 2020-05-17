package menu

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(menu *Menu) (*Menu, error)
	Update(menu *Menu) (*Menu, error)
	Delete(menu *Menu) error
	RemoveFood(menu *Menu) error
	FindAll()([]*Menu, error)
	FindById(menu *Menu)(*Menu, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(menu *Menu) (*Menu, error) {
	result, err := s.repo.Create(menu)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(menu *Menu) (*Menu, error) {
	result, err := s.repo.Update(menu)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) RemoveFood(menu *Menu) error {
	err := s.repo.RemoveFood(menu)
	if err != nil{
		return err
	}
	return nil
}

func (s service) Delete(menu *Menu) error {
	err := s.repo.Delete(menu)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Menu, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(menu *Menu) (*Menu, error) {
	result, err := s.repo.FindById(menu)
	if err != nil{
		return result, err
	}
	return result, nil
}
