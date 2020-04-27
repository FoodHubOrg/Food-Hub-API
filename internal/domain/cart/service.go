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
	FindById(cart *Cart)(*Cart, error)
	//CheckUser(cart *Cart) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//func (s service) CheckUser(cart *Cart) error {
//	rest, err := s.repo.FindById(cart.ID)
//	if err != nil {
//		return err
//	}
//
//	if rest.UserID != cart.UserID {
//		return fmt.Errorf("is not owner")
//	}
//
//	return nil
//}

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

func (s service) Delete(cart *Cart) error {
	//if err := s.CheckUser(cart); err != nil {
	//	return err
	//}
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

func (s service) FindById(cart *Cart) (*Cart, error) {
	result, err := s.repo.FindById(cart)
	if err != nil{
		return result, err
	}
	return result, nil
}
