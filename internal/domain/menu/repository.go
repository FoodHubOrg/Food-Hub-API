package menu

type Repository interface {
	Create(menu *Menu) (*Menu, error)
	Update(menu *Menu) (*Menu, error)
	Delete(menu *Menu) error
	FindAll()([]*Menu, error)
	FindById(menu *Menu)(*Menu, error)
	RemoveFood(menu *Menu) error
}
