package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/bedirhankilic/movieapicase/docs"

	"github.com/bedirhankilic/movieapicase/controller"
	dataaccesslayer "github.com/bedirhankilic/movieapicase/data-access-layer"
	"github.com/joho/godotenv"
)

// @title           Movie API Swagger
// @version         1.0
// @description     This is a movie api.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// @host      localhost:5000
// @BasePath  /
func main() {

	LoadEnv()
	InitDb()

	controller.InitHttpServer()
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func InitDb() {
	dbError := dataaccesslayer.InitDb()
	if dbError != nil {
		fmt.Println("DB Connection Error!")
		os.Exit(0)
	}
}
