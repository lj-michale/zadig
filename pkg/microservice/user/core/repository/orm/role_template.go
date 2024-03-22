package orm

import (
	"github.com/koderover/zadig/v2/pkg/microservice/user/core/repository/models"
	"gorm.io/gorm"
)

func CreateRoleTemplate(roleTemplate *models.RoleTemplate, db *gorm.DB) error {
	if err := db.Create(&roleTemplate).Error; err != nil {
		return err
	}
	return nil
}

func GetRoleTemplate(name string, db *gorm.DB) (*models.RoleTemplate, error) {
	resp := new(models.RoleTemplate)
	err := db.Where("name = ?", name).Find(&resp).Error

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ListRoleTemplates(db *gorm.DB) ([]*models.NewRole, error) {
	resp := make([]*models.NewRole, 0)

	err := db.Find(&resp).Error

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetRoleTemplateByID(id uint, db *gorm.DB) (*models.RoleTemplate, error) {
	resp := new(models.RoleTemplate)
	err := db.Where("id = ?", id).Find(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateRoleTemplateInfo(id uint, role *models.RoleTemplate, db *gorm.DB) error {
	err := db.Model(&models.RoleTemplate{}).Where("id = ?", id).Updates(role).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteRoleTemplateByName(name string, db *gorm.DB) error {
	var role models.RoleTemplate
	return db.Model(&models.RoleTemplate{}).
		Where("name = ?", name).
		Delete(&role).
		Error
}
