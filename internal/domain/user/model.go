package user

import (
	"Food-Hub-API/internal/database"
)

type User struct {
	database.Base
	Cpf string `gorm:"type:varchar(100);"`
	Name string `gorm:"type:varchar(100);not_null"`
	Email string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(250);not_null"`
}