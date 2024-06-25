package repo

import (
	"Pc_Build_Service/Internal/database/repo/model"
	model2 "Pc_Build_Service/Internal/userinterface/service/model"
	"context"
)

// Function to Create a Record of User in Db
func (m *impl) CreateUser(ctx context.Context, req *model2.UserRegistrationRequest) error {

	return m.db.
		Model(&model.UserRegistrationModel{}).
		Create(req).
		Error
}
