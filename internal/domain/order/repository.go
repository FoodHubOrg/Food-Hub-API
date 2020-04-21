package order

type Repository interface {
	Create(order *Order) (*Order, error)
	Update(order *Order, mode string) (*Order, error)
	Delete(order *Order) error
	FindAll()([]*Order, error)
	FindById(order *Order, mode string)(*Order, error)
}
