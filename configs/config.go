package configs

import "gorm.io/gorm"

var DB *gorm.DB

const (
	DBUser = "root"
	DBPass = "6233"
	DBName = "blockchain_gin"
	DBHost = "localhost"
	DBPort = "3306"
)
