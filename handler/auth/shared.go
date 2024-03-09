package auth

import (
	"ServerApp/handler/utils"
	"net/http"
)

var (
	ErrInvalidPayload     = "Cannot handle the request."
	ErrEmptyValue         = "All fields are required."
	ErrInvalidCredentails = "Username or Password incorrect."
	ErrInternalError      = "Server error, try again."
	ErrUnauthorized       = "Unauthorized access."
)

func (h *AuthHandler) UserStatus(w http.ResponseWriter, r *http.Request) {
	info, err := utils.AuthMiddleware(r)
	if err != nil {
		utils.SendErrorResponse(w, 400, ErrUnauthorized)
		return
	}
	user, err := h.svc.UserStatus(info.UserID)
	if err != nil {
		utils.SendErrorResponse(w, 400, ErrUnauthorized)
		return
	}
	utils.SendJSONResponse(w, 200, user)
}

func (h *AuthHandler) UserLogout(w http.ResponseWriter, r *http.Request) {
	if _, err := utils.AuthMiddleware(r); err != nil {
		utils.SendErrorResponse(w, 401, ErrUnauthorized)
		return
	}
	http.SetCookie(w, utils.NewEmptyCookie())
	utils.SendJSONResponse(w, http.StatusOK, nil)
}

func (h *AuthHandler) UserResetPwd(http.ResponseWriter, *http.Request) {}
