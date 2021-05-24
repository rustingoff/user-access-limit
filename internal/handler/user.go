package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rustingoff/user-access-limit/internal/models"
	"github.com/rustingoff/user-access-limit/internal/privileges"
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

func (h *Handler) testingRoute(w http.ResponseWriter, r *http.Request, privileg map[uint]bool) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Name":   "Iliaposmac",
		"Skills": "adsadsaddas",
	})
}
