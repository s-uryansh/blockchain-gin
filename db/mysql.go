package db

import (
	"fmt"
	"log"

	"github.com/s-uryansh/blockchain-gin/configs"
	"github.com/s-uryansh/blockchain-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	//creating database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configs.DBUser, configs.DBPass, configs.DBHost, configs.DBPort, configs.DBName)
	//opening database and checking if there is any error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	configs.DB = db
	log.Println("Database connected")

	err = configs.DB.AutoMigrate(&models.Block{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("database connected and migrated successfully")

}
