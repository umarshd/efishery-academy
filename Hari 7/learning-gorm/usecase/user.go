package usecase

import (
	"learning-gorm/entity"
	"learning-gorm/repository"
)

type IUserUsecase interface {
	CreateUser(user entity.UserRequest) (entity.User, error)
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepository}
}

func (useCase UserUseCase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{}

	u.Username = user.Username
	u.Email = user.Email
	u.Phone = user.Phone

	users, err := useCase.userRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}
	userRes := entity.UserResponse{
		ID:       users.ID,
		Username: user.Username,
		Email:    users.Email,
		Phone:    users.Phone,
	}

	return userRes, nil

}
