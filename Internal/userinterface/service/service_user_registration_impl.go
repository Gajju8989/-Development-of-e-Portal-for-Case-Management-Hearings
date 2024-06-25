package service

import (
	"Pc_Build_Service/Internal/genericresponse"
	"Pc_Build_Service/Internal/userinterface/service/model"
	"context"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (s *impl) UserRegistration(ctx context.Context,
	req *model.UserRegistrationRequest) error {
	// Check whether req is nil or not.
	if req == nil {
		return &genericresponse.GenericResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Req is nil",
		}
	}
	//Hash password before storing it.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &genericresponse.GenericResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to Hash Password",
		}
	}
	//Create userInformationPayload to store in db.
	userInformation := &model.UserRegistrationRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		MobileNo: req.MobileNo,
		Gender:   req.Gender,
	}
	//Store information in db.
	err = s.mysqlStore.CreateUser(ctx, userInformation)
	if err != nil {
		return &genericresponse.GenericResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable To Store User Information in Db",
		}
	}
	return nil
}
