package cart

type Repository interface {
	Create(cart *Cart) (*Cart, error)
	Update(cart *Cart) (*Cart, error)
	Delete(cart *Cart)  error
	FindAll()([]*Cart, error)
	FindById(cart *Cart)(*Cart, error)
}
