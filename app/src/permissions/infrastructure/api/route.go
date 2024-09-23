package api

import (
	"net/http"

	"github.com/melendez-dev/harvest-api-go/app/src/permissions/application"
	"github.com/melendez-dev/harvest-api-go/app/src/shared/controllers"
)

type PermissionHandler struct {
	services application.PermissionService
	controllers.BaseController
}

func Routes(service *application.PermissionService) http.Handler {
	mux := http.NewServeMux()
	pc := PermissionHandler{
		services: *service,
	}

	mux.HandleFunc("/permissions", pc.GetPermissions)

	return http.StripPrefix("/api/v1/rbac", mux)
}

func (h *PermissionHandler) GetPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, _ := h.services.FindAll()

	h.SendJson(w, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "success",
		"data":    permissions,
	})
}
