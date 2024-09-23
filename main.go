package main

import (
	"fmt"
	"net/http"

	"github.com/melendez-dev/harvest-api-go/app/config"

	"github.com/melendez-dev/harvest-api-go/app/src/bootstrap"
	"github.com/melendez-dev/harvest-api-go/app/src/permissions/infrastructure/api"
)

func main() {
	mux := http.NewServeMux()

	dbConn := bootstrap.InitDB("./app/config/config.yml")

	permissionsService := bootstrap.InitPermissions(dbConn)
	mux.Handle("/api/v1/rbac/permissions", api.Routes(permissionsService))

	conf, _ := config.NewConfig("./app/config/config.yml")
	if err := http.ListenAndServe(conf.Server.Port, mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
