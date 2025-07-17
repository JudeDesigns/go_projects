package database

import (
	"TodoManager/config"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB(cfg config.AppConfig) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

}
