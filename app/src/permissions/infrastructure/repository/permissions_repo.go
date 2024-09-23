package repository

import (
	"database/sql"
	"fmt"

	"github.com/melendez-dev/harvest-api-go/app/src/permissions/domain"
)

type PermissionsRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) *PermissionsRepository {
	return &PermissionsRepository{db: db}
}

func (p *PermissionsRepository) FindAll() ([]domain.Permissions, error) {
	rows, err := p.db.Query("SELECT id, uuid, code, name, status FROM permissions")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var permissions []domain.Permissions

	for rows.Next() {
		var permission domain.Permissions
		if err := rows.Scan(&permission.ID, &permission.UUID, &permission.Code, &permission.Name, &permission.Status); err != nil {
			fmt.Println(err)
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
