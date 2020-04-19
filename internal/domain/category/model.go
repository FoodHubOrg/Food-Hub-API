package category

import (
	"Food-Hub-API/internal/database"
)

type Category struct {
	database.Base
	Name  string `gorm:"not_null;unique_index"`
}
