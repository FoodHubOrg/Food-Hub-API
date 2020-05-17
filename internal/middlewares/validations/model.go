package validations

import (
	uuid "github.com/satori/go.uuid"
)


type User struct {
	ID uuid.UUID
	Email string `json:"email"`
	Password string `json:"password"`
	Name string `json:"name"`
	IsAdmin bool `json:"is_admin"`
	IsVerified bool `json:"is_verified"`
}

type Order struct {
	ID uuid.UUID
	Street string `gorm:"type:varchar(100);not_null"`
	Number string `gorm:"type:varchar(100);not_null"`
	City string `gorm:"type:varchar(100);not_null"`
	District string `gorm:"type:varchar(100);not_null"`
	Country string `gorm:"type:varchar(100);not_null"`
	PaymentType string `gorm:"type:varchar(100);not_null"`
	Status string `gorm:"type:varchar(100);default:'pending'"`
	CartID uuid.UUID `gorm:"type:uuid;not_null;unique"`
	UserID uuid.UUID `gorm:"type:uuid;not_null"`
	RestaurantID uuid.UUID `gorm:"type:uuid;not_null"`
	Foods []Food `gorm:"many2many:food_orders;"`
}

type Food struct {
	ID uuid.UUID
	Name string `json:"name"`
	Price string `json:"price"`
	MenuID uuid.UUID `gorm:"type:uuid;not_null"`
}

type Restaurant struct {
	ID uuid.UUID
	Name string `json:"name"`
	Location string `json:"location"`
	Cover string `json:"cover"`
	Time string `json:"time"`
	UserID  uuid.UUID `json:"user_id"`
	Categories []Category `json:"categories"`
	//Orders []Order
	//Menus []Menu
}

type Category struct {
	ID uuid.UUID
	Name  string `json:"name"`
	//Restaurants []Restaurant `gorm:"many2many:restaurant_categories;"`
}

type Menu struct {
	ID uuid.UUID
	Name string `json:"name"`
	RestaurantID uuid.UUID  `json:"restaurant_id"`
	UserID uuid.UUID `json:"user_id"`
	Foods []Food `json:"foods"`
}