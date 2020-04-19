package user

type Repository interface {
	CreateAccount(user *User) error
	FindById(id int) error
	GetUser(user *User) (*User, error)
	Update(user *User) error
}
