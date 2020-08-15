package main

import (
	models "models"
	routes "routes"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	models.InitDb()
	defer models.CloseDB()
	r := gin.Default()
	routes.SetupRoute(r)
	r.Run()
}
