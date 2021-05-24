package handler

import (
	"net/http"

	"github.com/rustingoff/user-access-limit/internal/service"
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
	router.HandleFunc("/test", TestingRoute("/", h.testingRoute))

	return router
}
