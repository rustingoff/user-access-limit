package handler

import (
	"net/http"
	"user_access_limit/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/view", CheckPrivileges("/", h.getUsers))

	return router
}