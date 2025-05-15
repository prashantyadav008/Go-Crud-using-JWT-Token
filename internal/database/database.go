package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-crud-msql/internal/config"
	"go-crud-msql/internal/models"
)

var DB *gorm.DB

func ConnectionDB(cfg *config.Config) {
	// dsn ----> Data Source Name (root:password@tcp(127.0.0.1:3306)/go_cruddb?parseTime=true)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = database
	log.Println("Database connected successfully!")

}

func RunMigrations(DB *gorm.DB) {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database Migrated Successfully!")
}
