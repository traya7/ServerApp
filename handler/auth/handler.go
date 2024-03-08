package auth

import (
	"AppServer/service"
	"github.com/gorilla/mux"
)

var (
	ErrInvalidPayload     = "Cannot handle the request."
	ErrEmptyValue         = "All fields are required."
	ErrInvalidCredentails = "Username or Password incorrect."
	ErrInternalError      = "Server error, try again."
	ErrUnauthorized       = "Unauthorized access."
)

type AuthHandler struct {
	svc *service.AuthService
}

func New() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Route(m *mux.Router) {
	// USER_HANDLERS
	m.HandleFunc("POST /api/v1/auth/login", h.ClientLogin)

	// ADMIN_HANDLERS
	m.HandleFunc("POST /api/v1/auth/u/login", h.AdminLogin)
	m.HandleFunc("POST /api/v1/auth/u/create", h.UserCreate)
	m.HandleFunc("POST /api/v1/auth/u/reset", h.UserReset)

	// USER+ADMIN
	m.HandleFunc("POST /api/v1/auth/status", h.UserStatus)
	m.HandleFunc("POST /api/v1/auth/resetpwd", h.UserResetPwd)
}
