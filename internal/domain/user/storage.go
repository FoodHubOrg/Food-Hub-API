package user

import (
	"github.com/jinzhu/gorm"
)

type Connection struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&User{})
	return &Connection{db}
}

func (c *Connection) CreateAccount(user *User) error {
	if err := c.db.Create(user).Error; err != nil{
		return err
	}
	return nil
}

func (c *Connection) FindById(id int) error {
	return nil
}

func (c *Connection) GetUser(user *User) error {
	if err := c.db.Where("email = ?", user.Email).First(&user).Error; err != nil{
		return err
	}
	return nil
}

func (c *Connection) Update(user *User) error {
	return nil
}