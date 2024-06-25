package service

import (
	"Pc_Build_Service/Internal/database/repo"
	"Pc_Build_Service/Internal/userinterface/service/model"
	"context"
)

type UserInterfaceService interface {
	UserRegistration(ctx context.Context,
		req *model.UserRegistrationRequest) error
	UserLogin(ctx context.Context, req *model.UserLoginRequest) (string, error)
}

type impl struct {
	mysqlStore repo.UserInterfaceRepo
}

func NewUserInterfaceSvc(mysqlStore repo.UserInterfaceRepo) UserInterfaceService {
	return &impl{
		mysqlStore: mysqlStore,
	}
}
