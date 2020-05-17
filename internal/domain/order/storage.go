package order

import (
	"fmt"
	"foodhub-api/internal/domain/cart"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

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
	//var foods []*food.Food
	var crt cart.Cart

	// fetch foods
	err := c.db.Set("gorm:auto_preload", true).Where("id = ?",
		order.CartID).First(&crt).Error
	if err != nil {
		return nil, err
	}

	logrus.Println(crt.Foods)

	if len(crt.Foods) < 1{
		return nil, fmt.Errorf("you cannot checkout an empty cart")
	}

	order.RestaurantID = crt.RestaurantID

	if err := c.db.Set("gorm:save_associations", false).Create(order).Error; err != nil {
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

	err := c.db.Set("gorm:auto_preload", true).Where("id = ?", order.ID).First(&order).Error
	if err != nil{
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
		err := c.db.Where("id = ?", order.ID).Preload("Foods").First(&order).Error
		if err != nil {
			return nil, err
		}
	case mode == "user":
		err := c.db.Where("user_id = ?", order.UserID).Preload("Foods").First(&order).Error
		if err != nil {
			return nil, err
		}
	default:
		return order, nil
	}

	return order, nil
}
