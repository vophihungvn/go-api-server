module hungvo.me/webapi/api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	models v0.0.0
	routes v0.0.0
// services v0.0.0
)

replace routes => ./routes

replace services => ./services

replace models => ./models
