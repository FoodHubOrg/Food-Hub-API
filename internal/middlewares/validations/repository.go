package validations

type Repository interface {
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(user *User)  error
	FindAll()([]*User, error)
	FindById(user *User)(*User, error)
	FindRestaurant(restaurant *Restaurant) error
	FindMenu(menu *Menu) error
}
