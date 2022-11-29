package routes

import (
	controller "github.com/shub-dab/Go-JWT-Authentication/controllers"
	"github.com/shub-dab/Go-JWT-Authentication/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
