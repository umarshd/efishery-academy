package handler

import (
	"learning-gorm/entity"
	"learning-gorm/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase}
}

func (handler UserHandler) CreateUser(c echo.Context) error {
	req := entity.UserRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := handler.userUseCase.CreateUser(req)

	if err != nil {
		return err
	}

	return c.JSON(201, user)
}
