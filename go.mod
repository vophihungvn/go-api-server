module hungvo.me/webapi/api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	routes v0.0.0
// services v0.0.0
)

replace routes => ./routes

replace services => ./services
