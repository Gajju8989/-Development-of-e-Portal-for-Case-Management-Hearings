package repo

import (
	"Pc_Build_Service/Internal/database/repo/model"
	model2 "Pc_Build_Service/Internal/userinterface/service/model"
	"context"
	"gorm.io/gorm"
)

type UserInterfaceRepo interface {
	// CreateUser inserts a new user registration record into the database.
	CreateUser(ctx context.Context, req *model2.UserRegistrationRequest) error
	GetUser(ctx context.Context, Email string) (*model.UserRegistrationModel, error)
}

type impl struct {
	db *gorm.DB
}

func NewMySqlStore(db *gorm.DB) UserInterfaceRepo {
	return &impl{db: db}
}
