package user

import (
	"aqrus/Microservice/gateway-api/utils"
)

type UserService interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	CreateUser(name string, email string, password string) (*User, error)
	UpdateUser(name string, email string, password string) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (userService *userService) GetAllUsers() ([]User, error) {
	return userService.repo.FindAll()
}
func (userService *userService) GetUserByID(id uint) (User, error) {
	return userService.repo.FindByID(id)

}

func (userService *userService) CreateUser(name string, email string, password string) (*User, error) {
	hashed_password, err := utils.HashPassWord(password)
	if err != nil {
		return nil, err
	}
	user := User{Name: name, Email: email, Password: hashed_password}
	return userService.repo.Save(user)
}
func (userService *userService) UpdateUser(name string, email string, password string) (*User, error) {
	hashed_password, err := utils.HashPassWord(password)
	if err != nil {
		return nil, err
	}
	user := User{Name: name, Email: email, Password: hashed_password}
	return userService.repo.Save(user)
}
