package user

import (
	"github.com/jinzhu/gorm"
)

type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &Connection{db}
}

func (c *Connection) Create(user *User) (*User, error) {
	if err := c.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Connection) Delete(user *User) error {
	err := c.db.Where("id = ?", user.ID).Delete(User{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c *Connection) FindBy(user *User, mode string) (*User, error) {
	switch {
	case mode == "email":
		if err := c.db.Where("email = ?", user.Email).Preload(
			"Orders").Preload("Restaurants").First(&user).Error; err != nil{
			return user, err
		}
	case mode == "id":
		if err := c.db.Where("id = ?", user.ID).Preload(
			"Orders").Preload("Restaurants").First(&user).Error; err != nil{
			return user, err
		}
	default:
		return user, nil
	}
	return user, nil
}

func (c *Connection) FindAll() ([]*User, error) {
	var users []*User
	if err := c.db.Preload("Orders").Preload("Restaurants").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (c *Connection) Update(user *User) (*User, error) {
	err := c.db.Model(user).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
