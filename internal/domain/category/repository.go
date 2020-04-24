package category

import (
	"Food-Hub-API/internal/domain/restaurant"
	uuid "github.com/satori/go.uuid"
)


type Repository interface {
	Create(category *restaurant.Category) (*restaurant.Category, error)
	Update(id uuid.UUID, category *restaurant.Category) (*restaurant.Category, error)
	Delete(id uuid.UUID) error
	FindAll()([]*restaurant.Category, error)
	FindById(id uuid.UUID)(*restaurant.Category, error)
}
