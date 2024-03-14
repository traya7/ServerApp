package auth

import (
	"ServerApp/handler/utils"
	"ServerApp/types"
	"encoding/json"
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
	user := r.Context().Value("data")
	utils.SendJSONResponse(w, 200, user)
}

func (h *AuthHandler) UserLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, utils.NewEmptyCookie())
	utils.SendJSONResponse(w, http.StatusOK, nil)
}

func (h *AuthHandler) UserCreate(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("data").(*types.User)
	if user.Role == "USER" {
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

	err := h.svc.CreateAccount(rd.Username, rd.Password, user.ID, user.Role)
	if err != nil {
		utils.SendErrorResponse(w, 400, ErrInternalError)
		return
	}
	utils.SendJSONResponse(w, http.StatusOK, nil)
}

func (h *AuthHandler) UserResetPwd(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("data").(*types.User)

	var rd struct {
		OldPwd string `json:"old_pwd"`
		NewPwd string `json:"new_pwd"`
	}

	if err := json.NewDecoder(r.Body).Decode(&rd); err != nil {
		utils.SendErrorResponse(w, 400, ErrInvalidPayload)
		return
	}

	if rd.OldPwd == "" || rd.NewPwd == "" {
		utils.SendErrorResponse(w, 400, ErrEmptyValue)
		return
	}

	if err := h.svc.UpdateMyPwd(user.ID, rd.OldPwd, rd.NewPwd); err != nil {
		utils.SendErrorResponse(w, 400, err.Error())
		return
	}

	utils.SendJSONResponse(w, 200, map[string]any{})
}

func (h *AuthHandler) UserResetClient(http.ResponseWriter, *http.Request) {}
