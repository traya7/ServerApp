package auth

import (
	"ServerApp/service"
	"github.com/gorilla/mux"
)

type AuthHandler struct {
	svc *service.AuthService
}

func New(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: s,
	}
}

func (h *AuthHandler) Route(m *mux.Router) {
	// USER_HANDLERS
	m.HandleFunc("/login", h.ClientLogin).Methods("POST")

	// ADMIN_HANDLERS
	m.HandleFunc("/u/login", h.AdminLogin).Methods("POST")
	m.HandleFunc("/u/create", h.UserCreate).Methods("POST")
	m.HandleFunc("/u/reset", h.UserReset).Methods("POST")

	// USER+ADMIN
	m.HandleFunc("/status", h.UserStatus).Methods("POST")
	m.HandleFunc("/resetpwd", h.UserResetPwd).Methods("POST")
}
