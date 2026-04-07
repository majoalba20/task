package main

import (
	"fmt"
	"go-repaso/internal/config"
	"go-repaso/internal/handler"
	"go-repaso/internal/queue"
	"go-repaso/internal/repository"
	"go-repaso/internal/routes"
	"go-repaso/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//repaso()
	setup()
}

func setup() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar .env, usando variables del sistema")
	}

	db := config.ConnectDB()

	taskRepo := repository.NewTaskRepository(db)

	taskQueue := queue.NewTaskQueue(100)
	taskWorker := queue.NewTaskWorker(taskRepo, taskQueue)
	taskWorker.Start()

	taskService := service.NewTaskService(taskRepo, taskQueue)
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

func repaso() {
	p := handler.Perro{Raza: "shit zu", Color: "negro"}
	g := handler.Gato{Edad: 2}

	handler.HacerSonido(p)
	handler.HacerSonido(g)
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.DataPerro(p)
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.DeclaracionVariables()
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.Punteros()
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.GoRoutines()
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.Channels()
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.ContextCancelado()
	handler.ContextTimeOut()
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.ManejoError()
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.Imprimir(45)
	handler.Imprimir("Hola mundo")
	fmt.Println("++++++++++++++++++++++++++++++")
	handler.Defer()
}
