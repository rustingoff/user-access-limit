package handler

import (
	"encoding/json"
	"net/http"
	"user_access_limit/internal/models"
	"user_access_limit/internal/privileges"
)

var user *models.User = &models.User{
	ID:         1,
	Name:       "Ion",
	Email:      "mail@mail.ru",
	Password:   "blya",
	Privileges: privileges.Admin,
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request, privileg map[uint]bool) {
	if !privileg[privileges.ViewDrivers] {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
	}
	json.NewEncoder(w).Encode(user)
}