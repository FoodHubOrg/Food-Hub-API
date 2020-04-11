package database

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	//"os"
)


func PostgresConnection() *gorm.DB {
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//dbName := os.Getenv("DB_NAME")
	//dbPassword := os.Getenv("DB_PASSWORD")
	//dbUsername := os.Getenv("DB_USERNAME")
	//dbMode := os.Getenv("DB_SSLMODE")
	//dbInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	dbHost, dbPort, dbUsername, dbName, dbPassword, dbMode)

	db, err := gorm.Open("postgres","postgres://postgres:kevina52@localhost:5432/foodhub?sslmode=disable")
	if err != nil{
		logrus.Fatal(err)
	}
	logrus.Info("Postgresql client connected")
	return db
}