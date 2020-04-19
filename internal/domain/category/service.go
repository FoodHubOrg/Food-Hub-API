package category

import uuid "github.com/satori/go.uuid"

type Service interface {
	Create(category *Category) (*Category, error)
	Update(id uuid.UUID, category *Category) (*Category, error)
	Delete(id uuid.UUID) error
	FindAll()([]*Category, error)
	FindById(id uuid.UUID)(*Category, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(category *Category) (*Category, error) {
	result, err := s.repo.Create(category)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(id uuid.UUID, category *Category) (*Category, error) {
	result, err := s.repo.Update(id, category)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Delete(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Category, error) {
	result, err := s.repo.FindAll()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) FindById(id uuid.UUID) (*Category, error) {
	result, err := s.repo.FindById(id)
	if err != nil {
		return result, err
	}
	return result, nil
}