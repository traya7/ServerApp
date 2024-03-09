package auth

import (
	"ServerApp/handler/utils"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	var Payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&Payload); err != nil {
		utils.SendErrorResponse(w, 401, ErrInvalidPayload)
		return
	}

	if Payload.Username == "" || Payload.Password == "" {
		utils.SendErrorResponse(w, 401, ErrEmptyValue)
		return
	}

	res, err := h.svc.UserLogin(Payload.Username, Payload.Password, "OTHER")
	if err != nil {
		utils.SendErrorResponse(w, 401, ErrInvalidCredentails)
		return
	}

	cookie, err := utils.NewCookie(res)
	if err != nil {
		utils.SendErrorResponse(w, 400, ErrInternalError)
		return
	}
	http.SetCookie(w, cookie)
	utils.SendJSONResponse(w, 200, res)
}

func (h *AuthHandler) UserCreate(w http.ResponseWriter, r *http.Request) {
	session, err := utils.AuthMiddleware(r)
	if err != nil {
		utils.SendErrorResponse(w, 401, ErrUnauthorized)
		return
	}
	if session.Role == "USER" {
		utils.SendErrorResponse(w, 401, ErrUnauthorized)
		return
	}
	var rd struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&rd); err != nil {
		utils.SendErrorResponse(w, 400, ErrInvalidPayload)
		return
	}

	if rd.Username == "" || rd.Password == "" {
		utils.SendErrorResponse(w, 400, ErrEmptyValue)
		return
	}

	err = h.svc.CreateAccount(rd.Username, rd.Password, session.UserID, session.Role)
	if err != nil {
		utils.SendErrorResponse(w, 400, ErrInternalError)
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, nil)
}

func (h *AuthHandler) UserReset(http.ResponseWriter, *http.Request) {}
