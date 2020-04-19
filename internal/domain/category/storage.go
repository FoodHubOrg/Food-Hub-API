package category

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	_ "github.com/sirupsen/logrus"
)

type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Category{})
	return &Connection{db,}
}

func (c *Connection) Create(category *Category) (*Category, error) {
	err := c.db.Create(category).Error
	if err != nil{
		return category, err
	}
	return category, nil
}

func (c *Connection) Update(id uuid.UUID, category *Category) (*Category, error) {
	err := c.db.Model(category).Where("id = ?", id).Update("name", category.Name).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c *Connection) Delete(id uuid.UUID) error {
	err := c.db.Where("id = ?", id).Delete(Category{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c *Connection) FindAll() ([]*Category, error) {
	var categories []*Category
	//var rest restaurant.Restaurant
	err := c.db.Find(&categories).Error
	if err != nil{
		return categories, err
	}
	return categories, nil
}

func (c *Connection) FindById(id uuid.UUID) (*Category, error) {
	var category Category
	err := c.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return &category, err
	}
	return &category, nil
}
