package validations

import (
	"fmt"
	"github.com/jinzhu/gorm"
	//"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
	//uuid"github.com/satori/go.uuid"
)
type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&User{})
	return &Connection{db,}
}

func (c Connection) Create(user *User) (*User, error) {
	// Insert into menu if not ready there
	if err := c.db.Create(user).Error; err != nil{
		return nil, err
	}
	return user, nil
}

func (c Connection) Update(user *User) (*User, error) {
	// Update User
	if err := c.db.Model(user).Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (c Connection) Delete(user *User) error {
	err := c.db.Where("id = ?", user.ID).Delete(User{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*User, error) {
	var users []*User
	err := c.db.Find(&users).Error
	if err != nil{
		return nil, err
	}
	return users, nil
}

func (c Connection) FindById(user *User) (*User, error) {
	err := c.db.Where("id = ?", user.ID).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (c Connection) FindRestaurant(restaurant *Restaurant) error {
	var rest Restaurant
	err := c.db.Where("id = ?", restaurant.ID).First(&rest).Error
	if err != nil{
		return err
	}

	if restaurant.UserID != rest.UserID {
		return fmt.Errorf("you don't have permission to update this resource")
	}

	return nil
}

func (c Connection) FindMenu(menu *Menu) error {
	var men Menu
	err := c.db.Where("id = ?", menu.ID).First(&men).Error
	if err != nil{
		return err
	}

	if menu.UserID != men.UserID {
		return fmt.Errorf("you don't have permission to update this resource")
	}

	return nil
}

