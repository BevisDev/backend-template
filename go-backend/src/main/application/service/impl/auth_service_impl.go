package impl

import (
	"github.com/BevisDev/backend-template/src/main/adapter/dto/request"
	"github.com/BevisDev/backend-template/src/main/application/service"
)

type AuthServiceImpl struct {
}

func (a *AuthServiceImpl) SignUp(dto request.SignUpDTO) {
	
}

func (a *AuthServiceImpl) SignIn(dto request.SignInDTO) {
}

func NewAuthServiceImpl() service.IAuthService {
	return &AuthServiceImpl{}
}
