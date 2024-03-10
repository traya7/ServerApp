package auth

import (
	"ServerApp/handler/utils"
	"ServerApp/service"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	svc     *service.AuthService
	useAuth utils.UseAuthFunc
}

func New(useAuth utils.UseAuthFunc, s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		svc:     s,
		useAuth: useAuth,
	}
}

func (h *AuthHandler) Route(m *mux.Router) {
	// USER_HANDLERS
	m.HandleFunc("/login", h.LoginAs("USER")).Methods("POST")
	m.HandleFunc("/u/login", h.LoginAs("OTHER")).Methods("POST")

	// ADMIN_HANDLERS
	m.HandleFunc("/status", h.useAuth(h.svc, h.UserStatus)).Methods("POST")
	m.HandleFunc("/resetpwd", h.useAuth(h.svc, h.UserResetPwd)).Methods("POST")
	m.HandleFunc("/create", h.useAuth(h.svc, h.UserCreate)).Methods("POST")

	m.HandleFunc("/reset-client", h.useAuth(h.svc, h.UserResetClient)).Methods("POST")
}
