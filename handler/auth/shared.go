package auth

import (
	"ServerApp/handler/utils"
	"net/http"
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

func (h *AuthHandler) UserResetPwd(http.ResponseWriter, *http.Request) {}
