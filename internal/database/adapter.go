package database

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"os"

	//"os"
)


func PostgresConnection() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil{
		logrus.Fatal(err)
	}
	logrus.Info("MySQL client connected")
	return db
}