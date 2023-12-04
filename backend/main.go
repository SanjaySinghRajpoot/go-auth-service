package main

import (
	"log"
	"os"

	"github.com/SanjaySinghRajpoot/backend/config"
	"github.com/SanjaySinghRajpoot/backend/controller"
	"github.com/SanjaySinghRajpoot/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	r := gin.Default()
	config.Connect()
	routes.UserRoute(r)

	admin := r.Group("/api/admin") // admin excess APIs
	admin.POST("/user", controller.PostUsers)
	admin.DELETE("/user", controller.DeleteUser)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
