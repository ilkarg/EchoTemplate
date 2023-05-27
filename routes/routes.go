package routes

import (
	controllers "api/controllers"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/confirmation-email", controllers.EmailConfirmation)

	v1 := e.Group("/api/v1")

	v1.POST("/logout", controllers.Logout)
	v1.POST("/login", controllers.Login)
	v1.POST("/registration", controllers.Registration)
}
