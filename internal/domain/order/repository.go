package order

type Repository interface {
	Create(order *Order) (*Order, error)
	Update(order *Order) (*Order, error)
	Delete(id string) (*Order, error)
	FindAll()([]*Order, error)
	FindById(id string)(*Order, error)
}
