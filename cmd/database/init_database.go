package database

import (
	"fmt"
	"os"

	"github.com/durotimicodes/xanda_task_R3_D3/helpers"
	"github.com/durotimicodes/xanda_task_R3_D3/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Initialize database
func InitDatabase() {
	err := godotenv.Load(".env")
	helpers.HandlerErr(err)

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helpers.HandlerErr(err)

	db.AutoMigrate(&models.Spaceship{}, &models.Armament{})
	DB = db
}
