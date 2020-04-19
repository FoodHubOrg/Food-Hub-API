package category

import uuid "github.com/satori/go.uuid"


type Repository interface {
	Create(category *Category) (*Category, error)
	Update(id uuid.UUID, category *Category) (*Category, error)
	Delete(id uuid.UUID) error
	FindAll()([]*Category, error)
	FindById(id uuid.UUID)(*Category, error)
}
