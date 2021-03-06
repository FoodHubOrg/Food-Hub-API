package user

import "foodhub-api/internal/helpers"

type Service interface {
	Create(user *User, mode string) (*User, error)
	FindBy(user *User, mode string) (*User, error)
	Login(user *User, password string) error
	Update(user *User, mode string) (*User, error)
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

func (u *service) Create(user *User, mode string) (*User, error) {
	hash, err := helpers.GenerateHash([]byte(user.Password))
	if err != nil{
		return nil, err
	}

	user.Password = hash

	result, err := u.repo.Create(user, mode)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func(u *service) Login(user *User, password string) error {
	entity, err := u.repo.FindBy(user, "email")
	if err != nil{
		return err
	}

	err = helpers.CompareHash(entity.Password, password)
	if err != nil{
		return err
	}

	return nil
}

func (u *service) Update(user *User, mode string) (*User, error) {
	result, err := u.repo.Update(user, mode)
	if err != nil{
		return nil, err
	}
	return result, nil
}

func (u *service) FindBy(user *User, mode string) (*User, error) {
	entity, err := u.repo.FindBy(user, mode)
	if err != nil{
		return nil, err
	}

	return entity, nil
}