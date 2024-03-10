package auth

import (
	"ServerApp/handler/utils"
	"encoding/json"
	"net/http"
)

func (h *AuthHandler) LoginAs(as string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		res, err := h.svc.UserLogin(Payload.Username, Payload.Password, as)
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
}
