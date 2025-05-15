package routes

import (
	"go-crud-msql/internal/handler"
	"go-crud-msql/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	api := router.Group("/api")

	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users", userHandler.GetAllUsers)
		api.GET("/users/:id", userHandler.GetUserById)
		api.PUT("/users/:id", userHandler.UpdateUser)
		api.DELETE("/users/:id", userHandler.DeleteUser)
		api.POST("/login", userHandler.LoginUser)
		api.POST("/logout", userHandler.LogoutUser)
		api.GET("/logout", userHandler.GetUserProfile)
		api.GET("/profile", middleware.AuthMiddleware(), userHandler.GetUserProfile) // add middleware

	}
}
