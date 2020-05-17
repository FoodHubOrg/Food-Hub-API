package validations

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(user *User) error
	FindAll()([]*User, error)
	FindById(user *User)(*User, error)
	CheckOwner(rest *Restaurant) error
	CheckMenuOwner(menu *Menu) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) CheckOwner(restaurant *Restaurant) error {
	err := s.repo.FindRestaurant(restaurant)
	if err != nil {
		return err
	}
	return nil
}

func (s service) CheckMenuOwner(menu *Menu) error {
	err := s.repo.FindMenu(menu)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Create(user *User) (*User, error) {
	result, err := s.repo.Create(user)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(user *User) (*User, error) {
	result, err := s.repo.Update(user)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(user *User) error {
	err := s.repo.Delete(user)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*User, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(user *User) (*User, error) {
	result, err := s.repo.FindById(user)
	if err != nil{
		return result, err
	}
	return result, nil
}
