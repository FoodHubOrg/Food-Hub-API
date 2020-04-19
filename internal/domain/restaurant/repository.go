package restaurant

import uuid "github.com/satori/go.uuid"

type Repository interface {
	Create(restaurant *Restaurant) (*Restaurant, error)
	Update(restaurant *Restaurant) (*Restaurant, error)
	Delete(restaurant *Restaurant) error
	FindAll()([]*Restaurant, error)
	FindById(id uuid.UUID)(Restaurant, error)
}
