package bootstrap

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/melendez-dev/harvest-api-go/app/config"
	"github.com/melendez-dev/harvest-api-go/app/src/permissions/application"
	"github.com/melendez-dev/harvest-api-go/app/src/permissions/infrastructure/repository"
)

func InitDB(routePath string) *sql.DB {
	conf, _ := config.NewConfig(routePath)
	sqlClient, err := sql.Open(conf.Database.DriverName, conf.Database.URL)

	if err != nil {
		panic(err.Error())
	}

	return sqlClient
}

func InitPermissions(dbConn *sql.DB) *application.PermissionService {
	permissionRepo := repository.NewPermissionRepository(dbConn)
	return &application.PermissionService{PermissionRepository: permissionRepo}
}
