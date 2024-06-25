package migrate

import (
	"Pc_Build_Service/Internal/database/repo/model"
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.UserRegistrationModel{}); err != nil {
		return err
	}
	return nil
}
