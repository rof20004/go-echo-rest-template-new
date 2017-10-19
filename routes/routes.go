package routes

import (
	"github.com/labstack/echo"
	"github.com/rof20004/go-echo-rest-template-new/modules/user"
)

// Server - server handler
var Server = echo.New()

// InitRoutes - initialize routes
func InitRoutes() {
	// Users routes
	Server.GET("/users", user.ListUsers)
	Server.POST("/users", user.InsertUser)
	Server.DELETE("/users/:id", user.DeleteUser)
}
