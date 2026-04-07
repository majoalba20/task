package main

import (
	"go-repaso/internal/config"
	"go-repaso/internal/handler"
	"go-repaso/internal/repository"
	"go-repaso/internal/routes"
	"go-repaso/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, usando variables del sistema")
	}

	db := config.ConnectDB()

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// Framework para Apis HTTP
	router := gin.Default()
	routes.SetupRoutes(router, taskHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor corriendo en http://localhost:%s", port)
	router.Run(":" + port)
}
