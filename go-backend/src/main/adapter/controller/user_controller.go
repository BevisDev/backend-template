package controller

import (
	"github.com/BevisDev/backend-template/src/main/adapter/dto/request"
	"github.com/BevisDev/backend-template/src/main/adapter/dto/response"
	"github.com/BevisDev/backend-template/src/main/application/service"
	"github.com/BevisDev/backend-template/src/main/common/consts"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) CreateRole(c *gin.Context) {
	var roleDTO request.RoleDTO
	if err := c.ShouldBindJSON(&roleDTO); err != nil {
		response.BadRequest(c, consts.InvalidRequest)
		return
	}
	uc.userService.CreateRole()
}

func (uc *UserController) AssignRole() {

}
