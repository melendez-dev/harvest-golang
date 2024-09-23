package repository

import (
	"github.com/melendez-dev/harvest-api-go/app/src/permissions/domain"
	"gorm.io/gorm"
)

type PermissionsRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionsRepository {
	return &PermissionsRepository{db: db}
}

func (p *PermissionsRepository) FindAll() ([]domain.Permissions, error) {
	var permissions []domain.Permissions
	p.db.Find(&permissions)

	return permissions, nil
}
