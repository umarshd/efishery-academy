package usecase

import (
	"fmt"
	"tugas/entity"
	"tugas/repository"
)

type IUserUseCase interface {
	CreateUser(user entity.UserRequest) (entity.User, error)
	GetAllUser() ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
	UpdateUser(userRequest entity.UserRequest, id int) (entity.User, error)
	DeleteUser(id int) error
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepository}
}

func (useCase UserUseCase) CreateUser(user entity.UserRequest) (entity.UserResponse, error) {
	u := entity.User{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	users, err := useCase.userRepository.Store(u)

	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:    users.ID,
		Name:  users.Name,
		Email: users.Email,
		Phone: users.Phone,
	}

	return userRes, nil

}

func (useCase UserUseCase) GetAllUser() ([]entity.UserResponse, error) {
	users, err := useCase.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	fmt.Println(users)

	userRes := []entity.UserResponse{}
	for _, user := range users {
		dataUser := entity.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		}

		userRes = append(userRes, dataUser)
	}

	return userRes, nil

}

func (useCase UserUseCase) GetUserById(id int) (entity.UserResponse, error) {
	user, err := useCase.userRepository.FindById(id)
	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	return userRes, nil
}

func (useCase UserUseCase) UpdateUser(userReq entity.UserRequest, id int) (entity.UserResponse, error) {
	user, err := useCase.userRepository.FindById(id)
	if err != nil {
		return entity.UserResponse{}, err
	}

	dataUser := entity.User{
		ID:    user.ID,
		Name:  userReq.Name,
		Email: userReq.Email,
		Phone: userReq.Phone,
	}

	user, err = useCase.userRepository.Update(dataUser)
	if err != nil {
		return entity.UserResponse{}, err
	}

	userRes := entity.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	}

	return userRes, nil
}

func (useCase UserUseCase) Delete(id int) error {
	_, err := useCase.userRepository.FindById(id)
	if err != nil {
		return err
	}

	message := useCase.userRepository.Delete(id)
	return message
}
