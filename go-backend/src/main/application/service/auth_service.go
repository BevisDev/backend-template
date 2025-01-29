package service

import "github.com/BevisDev/backend-template/src/main/adapter/dto/request"

type IAuthService interface {
	SignIn(dto request.SignInDTO)
	SignUp(dto request.SignUpDTO)
}
