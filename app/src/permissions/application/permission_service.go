package application

import (
	"github.com/melendez-dev/harvest-api-go/app/src/permissions/domain"
)

type PermissionService struct {
	PermissionRepository domain.PermissionsInterface
}

func (s *PermissionService) FindAll() ([]domain.Permissions, error) {
	return s.PermissionRepository.FindAll()
}
