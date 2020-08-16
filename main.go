package main

import (
	"log"
	models "models"
	routes "routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitDb()
	defer models.CloseDB()
	r := gin.Default()
	routes.SetupRoute(r)
	r.Run()
}
