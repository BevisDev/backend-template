package impl

import "github.com/BevisDev/backend-template/src/main/application/service"

type UserServiceImpl struct {
}

func (u *UserServiceImpl) CreateUser() {
}

func (u *UserServiceImpl) CreateRole() {
}

func (u *UserServiceImpl) AssignRole() {
}

func NewUserServiceImpl() service.IUserService {
	return &UserServiceImpl{}
}
