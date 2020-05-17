package cart

import (
	"foodhub-api/internal/domain/food"
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
	if err := c.db.Set("gorm:save_associations",
		false).Where(Cart{UserID:cart.UserID,
			RestaurantID:cart.RestaurantID}).FirstOrCreate(cart).Error; err != nil{
		return nil, err
	}

	return cart, nil
}

func (c Connection) Update(cart *Cart) (*Cart, error) {
	// Update Cart
	var item food.Food

	// Insert into foods if not ready there
	err := c.db.Where("id = ?", cart.Foods[0].ID).First(&item).Error
	if err != nil {
		return nil, err
	}

	if err := c.db.Model(cart).Association(
		"Foods").Append(item).Error; err != nil {
		return nil, err
	}

	if err := c.db.Set("gorm:auto_preload", true).First(cart).Error; err != nil{
		return nil, err
	}

	return cart, nil
}

func (c Connection) Remove(cart *Cart) error {
	// Update Cart
	var item food.Food

	// Insert into categories if not ready there
	err := c.db.Where("id = ?", cart.Foods[0].ID).First(&item).Error
	if err != nil {
		return err
	}

	if err := c.db.Model(cart).Association(
		"Foods").Delete(item).Error; err != nil {
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
	err := c.db.Set("gorm:auto_preload", true).Find(&carts).Error
	if err != nil{
		return nil, err
	}
	return carts, nil
}

func (c Connection) FindByID(cart *Cart) (*Cart, error) {
	err := c.db.Set("gorm:auto_preload", true).Where("restaurant_id = ?", cart.RestaurantID).First(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}

