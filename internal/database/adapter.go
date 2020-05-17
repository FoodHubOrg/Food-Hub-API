package database

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"os"

	//"os"
)


func PostgresConnection() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil{
		logrus.Fatal(err)
	}
	logrus.Info("Postgres client connected")
	return db
}