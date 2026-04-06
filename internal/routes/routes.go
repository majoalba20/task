package routes

import (
	"go-repaso/internal/handler"
	"go-repaso/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", handler.Me)
	}
}
