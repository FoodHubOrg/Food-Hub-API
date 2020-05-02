package category

import (
	"food-hub-api/internal/domain/restaurant"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	_ "github.com/sirupsen/logrus"
)

type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &Connection{db,}
}

func (c *Connection) Create(category *restaurant.Category) (*restaurant.Category, error) {
	err := c.db.Create(category).Error
	if err != nil{
		return category, err
	}
	return category, nil
}

func (c *Connection) Update(id uuid.UUID, category *restaurant.Category) (*restaurant.Category, error) {
	err := c.db.Model(category).Where("id = ?", id).Update("name", category.Name).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *Connection) Delete(id uuid.UUID) error {
	err := c.db.Where("id = ?", id).Delete(restaurant.Category{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c *Connection) FindAll() ([]*restaurant.Category, error) {
	var categories []*restaurant.Category
	//var rest restaurant.Restaurant
	err := c.db.Preload("Restaurants").Find(&categories).Error
	if err != nil{
		return categories, err
	}
	return categories, nil
}

func (c *Connection) FindById(id uuid.UUID) (*restaurant.Category, error) {
	var category restaurant.Category
	err := c.db.Where("id = ?", id).Preload("Restaurants").First(&category).Error
	if err != nil {
		return &category, err
	}
	return &category, nil
}
