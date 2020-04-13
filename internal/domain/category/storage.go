package restaurant

import "github.com/jinzhu/gorm"

type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Restaurant{})
	return &Connection{db,}
}

func (c Connection) Create(restaurant *Restaurant) (*Restaurant, error) {
	panic("implement me")
}

func (c Connection) Update(restaurant *Restaurant) (*Restaurant, error) {
	panic("implement me")
}

func (c Connection) Delete(id string) (*Restaurant, error) {
	panic("implement me")
}

func (c Connection) FindAll() ([]*Restaurant, error) {
	panic("implement me")
}

func (c Connection) FindById(id string) (*Restaurant, error) {
	panic("implement me")
}
