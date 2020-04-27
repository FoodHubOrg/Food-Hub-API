package cart

import (
	"github.com/jinzhu/gorm"
	//"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
	//uuid"github.com/satori/go.uuid"
)
type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Cart{})
	return &Connection{db,}
}

func (c Connection) Create(cart *Cart) (*Cart, error) {
	// Insert into menu if not ready there
	if err := c.db.Create(cart).Error; err != nil{
		return nil, err
	}
	return cart, nil
}

func (c Connection) Update(cart *Cart) (*Cart, error) {
	// Update Cart
	if err := c.db.Model(cart).Updates(cart).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func (c Connection) Delete(cart *Cart) error {
	err := c.db.Where("id = ?", cart.ID).Delete(Cart{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Cart, error) {
	var carts []*Cart
	err := c.db.Find(&carts).Error
	if err != nil{
		return nil, err
	}
	return carts, nil
}

func (c Connection) FindById(cart *Cart) (*Cart, error) {
	err := c.db.Where("id = ?", cart.ID).First(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

