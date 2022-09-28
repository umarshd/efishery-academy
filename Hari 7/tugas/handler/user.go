package handler

import (
	"net/http"
	"strconv"
	"tugas/entity"
	"tugas/usecase"

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
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Create user failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "User created successfully",
		Data:    user,
	})
}

func (handler UserHandler) GetAllUser(c echo.Context) error {
	users, err := handler.userUseCase.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get user failed",
			Error:   err.Error(),
		})
	}

	if len(users) < 1 {
		return c.JSON(http.StatusNotFound, entity.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Get users Failed. No users found",
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Users found",
		Data:    users,
	})
}

func (handler UserHandler) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userUseCase.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get user Failed",
			Error:   err.Error(),
		})
	}

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, entity.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Get user Failed. No user found",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User found",
		Data:    user,
	})
}

func (handler UserHandler) UpdateUser(c echo.Context) error {
	req := entity.UserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Request Body",
			Error:   err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := handler.userUseCase.UpdateUser(req, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update user failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User updated successfully",
		Data:    user,
	})
}

func (handler UserHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := handler.userUseCase.Delete(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete user failed",
			Error:   err.Error(),
		}))
	}
	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User deleted successfully",
	})
}
