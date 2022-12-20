package service

import (
	"puppy/model"
	userViewModel "puppy/models/user"
	"puppy/repository"
)

type UserService interface {
	GetUserList() ([]model.User, error)
	CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error)
	GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (model.User, error)
}

type userService struct {
}

func NewUserService() userService {
	return userService{}
}

func (userService) GetUserList() ([]model.User, error) {

	userRepository := repository.NewUserRepository()
	userList, err := userRepository.GetUserList()

	return userList, err
}

func (userService) GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (model.User, error) {

	userRepository := repository.NewUserRepository()
	user, err := userRepository.GetUserByUserNameAndPassword(loginViewModel.UserName, loginViewModel.Password)

	return user, err
}

func (userService) CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error) {
	userEntity := model.User{
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
		BirthDate: userInput.BirthDate,
		UserName:  userInput.UserName,
		Password:  userInput.Password,
		Code:      userInput.Code,
	}

	userRepository := repository.NewUserRepository()
	userId, err := userRepository.InsertUser(userEntity)

	return userId, err
}
