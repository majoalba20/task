package routes

import (
	"go-repaso/internal/handler"
	"go-repaso/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, taskHandler *handler.TaskHandler) {
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/tasks", taskHandler.CreateTask)
		protected.GET("/tasks", taskHandler.GetTasks)
		protected.GET("/tasks/:id", taskHandler.GetTaskByID)
		protected.PATCH("/tasks/:id", taskHandler.UpdateTask)
		protected.DELETE("/tasks/:id", taskHandler.DeleteTask)
		protected.POST("/tasks/:id/process", taskHandler.ProcessTask)
	}
}
