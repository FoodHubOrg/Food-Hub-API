package order

import (
	"Food-Hub-API/internal/domain/cart"
	"github.com/jinzhu/gorm"
	//"github.com/sirupsen/logrus"
	//"github.com/sirupsen/logrus"
)

type Connection struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&Order{})
	return &Connection{db,}
}

func (c Connection) Create(order *Order) (*Order, error) {
	// declare
	//var foods []*food.Food
	var crt cart.Cart

	// Insert order first
	if err := c.db.Set("gorm:save_associations", false).Create(order).Error; err != nil {
		return nil, err
	}

	crt.ID = order.CartID

	// fetch foods
	err := c.db.Where("id = ?", crt.ID).Preload("Foods").First(&crt).Error
	if err != nil {
		return nil, err
	}

	// create association with orders
	if err := c.db.Model(order).Association("Foods").Append(crt.Foods).Error; err != nil {
		return nil, err
	}

	//// clear cart
	if err := c.db.Model(crt).Association("Foods").Clear().Error; err != nil {
		return nil, err
	}


	return order, nil
}

func (c Connection) Update(order *Order, mode string) (*Order, error) {
	// Update Order
	if err := c.db.Model(order).Updates(&Order{Status: mode}).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (c Connection) Delete(order *Order) error {
	err := c.db.Where("id = ?", order.ID).Delete(Order{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Order, error) {
	var orders []*Order
	err := c.db.Preload("Foods").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c Connection) FindById(order *Order, mode string) (*Order, error) {
	switch {
	case mode == "order":
		err := c.db.Where("id = ?", order.ID).First(order).Error
		if err != nil {
			return order, err
		}
	case mode == "user":
		err := c.db.Where("user_id = ?", order.UserID).First(order).Error
		if err != nil {
			return order, err
		}
	default:
		return order, nil
	}

	return order, nil
}
