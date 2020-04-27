package food

import (
	"Food-Hub-API/internal/database"
	uuid "github.com/satori/go.uuid"
)

type Food struct {
	database.Base
	Name string `gorm:"type:varchar(100);"`
	Price string `gorm:"type:varchar(100);"`
	MenuID uuid.UUID `gorm:"type:uuid;not_null"`
}
