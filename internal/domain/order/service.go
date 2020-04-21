package order

type Service interface {
	Create(order *Order) (*Order, error)
	Update(order *Order, mode string) (*Order, error)
	Delete(order *Order) error
	FindAll()([]*Order, error)
	FindById(order *Order, mode string)(*Order, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(order *Order) (*Order, error) {
	result, err := s.repo.Create(order)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s service) Update(order *Order, mode string) (*Order, error) {
	result, err := s.repo.Update(order, mode)
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) Delete(order *Order) error {
	err := s.repo.Delete(order)
	if err != nil{
		return err
	}
	return nil
}

func (s service) FindAll() ([]*Order, error) {
	result, err := s.repo.FindAll()
	if err != nil{
		return result, err
	}
	return result, nil
}

func (s service) FindById(order *Order, mode string) (*Order, error) {
	result, err := s.repo.FindById(order, mode)
	if err != nil{
		return result, err
	}
	return result, nil
}