package routes

import (
	"tugas/handler"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.GET("/users", userHandler.GetAllUser)
	e.GET("/users/:id", userHandler.GetUserById)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

}
