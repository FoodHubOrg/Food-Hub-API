package restaurant

import (
	"github.com/jinzhu/gorm"
	//"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
)
import uuid "github.com/satori/go.uuid"

type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Restaurant{})
	return &Connection{db,}
}

func (c Connection) Create(restaurant *Restaurant) (*Restaurant, error) {
	// Insert restaurant first
	if err := c.db.Set("gorm:save_associations",
		false).Create(restaurant).Error; err != nil {
		return nil, err
	}

	// Insert into categories if not ready there
	for i := range restaurant.Categories {
		if c.db.Where("name = ?",
			restaurant.Categories[i].Name).First(&restaurant.Categories[i]).RecordNotFound() {
			if err := c.db.Create(&restaurant.Categories[i]).Error; err != nil {
				return nil, err
			}
		}
	}

	// create association for categories
	if err := c.db.Model(restaurant).Association(
		"Categories").Append(restaurant.Categories).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (c Connection) Update(restaurant *Restaurant) (*Restaurant, error) {
	// Update Restaurant
	if err := c.db.Model(restaurant).Updates(&Restaurant{Name: restaurant.Name,
		Location:restaurant.Location, Categories:restaurant.Categories}).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (c Connection) Delete(restaurant *Restaurant) error {
	err := c.db.Where("id = ?", restaurant.ID).Delete(Restaurant{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Restaurant, error) {
	var restaurants []*Restaurant
	err := c.db.Preload("Categories").Find(&restaurants).Error
	if err != nil{
		return nil, err
	}
	return restaurants, nil
}

func (c Connection) FindById(id uuid.UUID) (Restaurant, error) {
	var restaurant Restaurant
	err := c.db.Where("id = ?", id).First(&restaurant).Error
	if err != nil {
		return restaurant, err
	}
	return restaurant, nil
}
