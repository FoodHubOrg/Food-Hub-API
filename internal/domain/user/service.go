package user

import "Food-Hub-API/internal/helpers"

type Service interface {
	CreateAccount(user *User) error
	Login(user *User, password string) error
	Update(user *User) error
}

type service struct {
	repo Repository
}

// Implement UserHandler Interface
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (u *service) CreateAccount(user *User) error {
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

func(u *service) Login(user *User, password string) error{
	entity, err := u.repo.GetUser(user)
	if err != nil{
		return err
	}

	err = helpers.CompareHash(entity.Password, password)
	if err != nil{
		return err
	}

	return nil
}

func (u *service) Update(user *User) error {
	err := u.repo.Update(user)
	if err != nil{
		return err
	}
	return nil
}
