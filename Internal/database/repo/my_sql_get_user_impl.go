package repo

import (
	"Pc_Build_Service/Internal/database/repo/model"
	"context"
)

func (s *impl) GetUser(ctx context.Context, Email string) (*model.UserRegistrationModel, error) {
	user := &model.UserRegistrationModel{}
	err := s.db.
		Model(&model.UserRegistrationModel{}).
		Where("email = ?", Email).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
