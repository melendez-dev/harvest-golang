package bootstrap

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/melendez-dev/harvest-api-go/app/config"
	"github.com/melendez-dev/harvest-api-go/app/src/permissions/application"
	"github.com/melendez-dev/harvest-api-go/app/src/permissions/infrastructure/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(routePath string) *gorm.DB {
	conf, _ := config.NewConfig(routePath)
	db, err := gorm.Open(mysql.Open(conf.Database.URL), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

func InitPermissions(dbConn *gorm.DB) *application.PermissionService {
	permissionRepo := repository.NewPermissionRepository(dbConn)
	return &application.PermissionService{PermissionRepository: permissionRepo}
}
