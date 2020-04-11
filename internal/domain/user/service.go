package user

import "Food-Hub-API/internal/helpers"

type UserService interface {
	CreateAccount(user *User) error
	Login(user *User, password string) error
}

type userService struct {
	repo UserRepository
}

// Implement UserHandler Interface
func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository,
	}
}

func (u *userService) CreateAccount(user *User) error {
	hash, err := helpers.GenerateHash([]byte(user.Password))
	if err != nil{
		return err
	}
	user.Password = hash

	err = u.repo.CreateAccount(user)
	if err != nil{
		return err
	}
	return nil
}

func(u *userService) Login(user *User, password string) error{
	err := u.repo.GetUser(user)
	if err != nil{
		return err
	}

	err = helpers.CompareHash(user.Password, password)
	if err != nil{
		return err
	}

	return nil
}
