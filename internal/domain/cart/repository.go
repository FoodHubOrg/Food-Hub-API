package cart

type Repository interface {
	Create(cart *Cart) (*Cart, error)
	Update(cart *Cart) (*Cart, error)
	Remove(cart *Cart) error
	Delete(cart *Cart)  error
	FindAll()([]*Cart, error)
	FindByID(cart *Cart)(*Cart, error)
}
