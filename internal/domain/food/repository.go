package food

type Repository interface {
	Create(food *Food) (*Food, error)
	Update(food *Food) (*Food, error)
	Delete(food *Food)  error
	FindAll()([]*Food, error)
	FindById(food *Food)(*Food, error)
}
