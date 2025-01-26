package configs

import "gorm.io/gorm"

var DB *gorm.DB

const (
	DBUser = "USER"
	DBPass = "DB_PASS"
	DBName = "DB_NAME"
	DBHost = "localhost"
	DBPort = "DB_PORT"
)
