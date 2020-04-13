package restaurant

type CategoryRepository interface {
	Create(category *Category) (*Category, error)
	Update(category *Category) (*Category, error)
	Delete(id string) (*Category, error)
	FindAll()([]*Category, error)
	FindById(id string)(*Category, error)
}
