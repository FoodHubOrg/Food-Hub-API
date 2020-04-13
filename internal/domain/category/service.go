package restaurant

type Service interface {
	Create(restaurant *Restaurant) (*Restaurant, error)
	Update(restaurant *Restaurant) (*Restaurant, error)
	Delete(id string) (*Restaurant, error)
	FindAll()([]*Restaurant, error)
	FindById(id string)(*Restaurant, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s service) Create(restaurant *Restaurant) (*Restaurant, error) {
	panic("implement me")
}

func (s service) Update(restaurant *Restaurant) (*Restaurant, error) {
	panic("implement me")
}

func (s service) Delete(id string) (*Restaurant, error) {
	panic("implement me")
}

func (s service) FindAll() ([]*Restaurant, error) {
	panic("implement me")
}

func (s service) FindById(id string) (*Restaurant, error) {
	panic("implement me")
}