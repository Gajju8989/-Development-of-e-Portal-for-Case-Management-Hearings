package handler

import (
	"Pc_Build_Service/Internal/userinterface/service"
	"github.com/gin-gonic/gin"
)

type UserInterfaceHandler interface {
	CreateUserHandler(c *gin.Context)
	UserLoginHandler(c *gin.Context)
}

type impl struct {
	userInterfaceSvc service.UserInterfaceService
}

func NewUserInterfaceHandler(userInterfaceSvc service.UserInterfaceService) UserInterfaceHandler {
	return &impl{
		userInterfaceSvc: userInterfaceSvc,
	}
}
