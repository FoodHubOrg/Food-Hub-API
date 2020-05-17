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

func (c *Connection) Create(user *User, mode string) (*User, error) {
	switch  {
	case mode == "social":
		if err := c.db.Where(User{Email: user.Email}).FirstOrCreate(&user).Error; err != nil {
			return nil, err
		}
	case mode == "system":
		if err := c.db.Create(&user).Error; err != nil {
			return nil, err
		}
	default:
		return user, nil
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
		if err := c.db.Set("gorm:auto_preload", true).Where(
			"email = ?", user.Email).First(&user).Error; err != nil{
			return user, err
		}
	case mode == "id":
		if err :=  c.db.Set("gorm:auto_preload", true).Where(
			"id = ?", user.ID).First(&user).Error; err != nil{
			return user, err
		}
	default:
		return user, nil
	}
	return user, nil
}

func (c *Connection) FindAll() ([]*User, error) {
	var users []*User
	if err := c.db.Set("gorm:auto_preload", true).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (c *Connection) Update(user *User, mode string) (*User, error) {
	switch {
	case mode == "makeRestaurantOwner":
		err := c.db.Model(&user).Update("is_restaurant_owner", true).Error
		if err != nil {
			return nil, err
		}
	case mode == "revokeRestaurantOwner":
		err := c.db.Model(&user).Update("is_restaurant_owner", false).Error
		if err != nil {
			return nil, err
		}
	default:
		return nil, nil
	}

	if err := c.db.Set("gorm:auto_preload", true).First(&user).Error; err != nil{
		return user, nil
	}

	return user, nil
}
