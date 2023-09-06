package config

import (
	"fmt"
	"rest-api-note-taking/models"
	"rest-api-note-taking/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := utils.GetEnv("DATABASE_USERNAME", "root")
	password := utils.GetEnv("DATABASE_PASSWORD", "")
	host := utils.GetEnv("DATABASE_HOST", "127.0.0.1")
	port := utils.GetEnv("DATABASE_PORT", "3306")
	database := utils.GetEnv("DATABASE_NAME", "db_note_taking")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{}, &models.Note{})
	return db
}
