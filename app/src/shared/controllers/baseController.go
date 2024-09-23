package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type BaseController struct {
	DB *sql.DB
}

func (bc *BaseController) SendJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, "Error encondig JSON", http.StatusInternalServerError)
	}
}
