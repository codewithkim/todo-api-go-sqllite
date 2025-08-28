package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"todo-api/config"
	"todo-api/routes"
)

func main() {
	fmt.Println("START TODO API")
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("😱 Panic: %v", r)
		}
	}()

	log.Println("Starting server...")
	godotenv.Load()

	fmt.Println("🔎 ENV check:")
	fmt.Println("  DB_PATH:", os.Getenv("DB_PATH"))

	config.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running at http://localhost:" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
