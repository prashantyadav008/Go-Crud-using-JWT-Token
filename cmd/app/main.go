package main

import (
	"go-crud-msql/internal/config"
	"go-crud-msql/internal/database"
	"go-crud-msql/internal/handler"
	"go-crud-msql/internal/repository"
	"go-crud-msql/internal/routes"
	"go-crud-msql/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Step 1: Load config
	cfg := config.LoadConfig()

	// Step 2: Connect to DB
	database.ConnectionDB(cfg)

	// Step 3: Auto migrate the model
	database.RunMigrations(database.DB)

	log.Println("✅ Application started successfully!")

	// Step 4: Create layers
	userRepo := repository.NewUserRepository(database.DB)
	log.Println("✅ Call repository!", userRepo)

	userService := service.NewUserService(userRepo)
	log.Println("✅ Call service!", userService)

	userHandler := handler.NewUserHandler(userService, database.DB)
	log.Println("✅ Call handler!", userService)

	// Step 6: Setup Gin router
	router := gin.Default()

	// Step 7: Setup routes
	routes.SetupRoutes(router, userHandler)

	// Step 8: Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
