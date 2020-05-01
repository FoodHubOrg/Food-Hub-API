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
	if err := c.db.Set("gorm:save_associations", false).Create(cart).Error; err != nil{
		return nil, err
	}

	if err := c.db.Model(cart).Association(
		"Foods").Append(cart.Foods).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (c Connection) Update(cart *Cart) (*Cart, error) {
	// Update Cart
	if err := c.db.Model(cart).Association(
		"Foods").Append(cart.Foods[0]).Error; err != nil {
		return nil, err
	}

	return cart, nil
}

func (c Connection) Remove(cart *Cart) error {
	if err := c.db.Model(cart).Association(
		"Foods").Delete(cart.Foods[0]).Error; err != nil {
		return err
	}

	return nil
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
	err := c.db.Preload("Foods").Find(&carts).Error
	if err != nil{
		return nil, err
	}
	return carts, nil
}

func (c Connection) FindByID(cart *Cart) (*Cart, error) {
	err := c.db.Preload("Foods").Where("restaurant_id = ?", cart.RestaurantID).First(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

