module hungvo.me/webapi/api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	models v0.0.0
	routes v0.0.0
// services v0.0.0
)

replace routes => ./src/routes

replace services => ./src/services

replace models => ./src/models
