package menu

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
	db.AutoMigrate(&Menu{})
	return &Connection{db,}
}

func (c Connection) Create(menu *Menu) (*Menu, error) {
	// Insert menu first
	if err := c.db.Set("gorm:save_associations",
		false).Create(menu).Error; err != nil {
		return nil, err
	}

	// Insert into categories if not ready there
	for i := range menu.Foods {
		if c.db.Where("name = ?",
			menu.Foods[i].Name).First(&menu.Foods[i]).RecordNotFound() {
			if err := c.db.Create(&menu.Foods[i]).Error; err != nil {
				return nil, err
			}
		}
	}

	// create association
	if err := c.db.Model(menu).Association(
		"Foods").Append(menu.Foods).Error; err != nil {
		return nil, err
	}

	return menu, nil
}

func (c Connection) Update(menu *Menu) (*Menu, error) {
	// Update Menu
	if err := c.db.Model(menu).Updates(&Menu{Foods:menu.Foods}).Error; err != nil {
		return nil, err
	}

	return menu, nil
}

func (c Connection) Delete(menu *Menu) error {
	err := c.db.Where("id = ?", menu.ID).Delete(Menu{}).Error
	if err != nil {
		return err
	}
	return err
}

func (c Connection) FindAll() ([]*Menu, error) {
	var menus []*Menu
	err := c.db.Preload("Foods").Find(&menus).Error
	if err != nil{
		return nil, err
	}
	return menus, nil
}

func (c Connection) FindById(menu *Menu) (*Menu, error) {
	err := c.db.Where("id = ?", menu.ID).First(&menu).Error
	if err != nil {
		return menu, err
	}
	return menu, nil
}

