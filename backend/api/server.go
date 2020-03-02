package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"todo-app/api/controllers"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}
}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println(".env was loaded!")
	}

	server.Initialize(
		os.Getenv("API_DB_DRIVER"),
		os.Getenv("MONGO_ATLAS_DB_USER"),
		os.Getenv("MONGO_ATLAS_DB_PASS"),
		os.Getenv("MONGO_ATLAS_DB_NAME"),
	)

	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	fmt.Println(fmt.Sprintf("Listening to port %s", apiPort))

	server.Run(apiPort)
}
