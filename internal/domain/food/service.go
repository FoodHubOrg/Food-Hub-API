package food

import (
	//"fmt"
	_ "github.com/satori/go.uuid"
)

type Service interface {
	Create(food *Food) (*Food, error)
	Update(food *Food) (*Food, error)
	Delete(food *Food) error
	FindAll()([]*Food, error)
	FindById(food *Food)(*Food, error)
	//CheckUser(food *Food) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//func (s service) CheckUser(food *Food) error {
//	rest, err := s.repo.FindById(food.ID)
//	if err != nil {
//		return err
//	}
//
//	if rest.UserID != food.UserID {
//		return fmt.Errorf("is not owner")
//	}
//
//	return nil
//}

func (s service) Create(food *Food) (*Food, error) {
	result, err := s.repo.Create(food)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(food *Food) (*Food, error) {
	result, err := s.repo.Update(food)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(food *Food) error {
	//if err := s.CheckUser(food); err != nil {
	//	return err
	//}
	err := s.repo.Delete(food)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Food, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(food *Food) (*Food, error) {
	result, err := s.repo.FindById(food)
	if err != nil{
		return result, err
	}
	return result, nil
}
