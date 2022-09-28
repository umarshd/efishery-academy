package main

import (
	"tugas/config"
	"tugas/handler"
	"tugas/repository"
	"tugas/routes"
	"tugas/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()

	config.Migrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	routes.Routes(e, userHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
