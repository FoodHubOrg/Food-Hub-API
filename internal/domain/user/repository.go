package user

type UserRepository interface {
	CreateAccount(user *User) error
	FindById(id int) error
	GetUser(user *User) error
	Update(user *User) error
}
