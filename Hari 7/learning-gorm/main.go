package main

import (
	"learning-gorm/config"
	"learning-gorm/handler"
	"learning-gorm/repository"
	"learning-gorm/routes"
	"learning-gorm/usecase"

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
