package handler

import (
	"ServerApp/handler/auth"
	"ServerApp/handler/utils"
	"ServerApp/handler/wallet"
	"ServerApp/service"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type Params struct {
	Auth *service.AuthService
}

func NewRouter(p Params) http.Handler {
	m := mux.NewRouter()

	auth.New(useAuth, p.Auth).Route(m.PathPrefix("/v1/auth").Subrouter())
	wallet.New().Route(m.PathPrefix("/v1/wallet").Subrouter())

	m.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Opps, not found page."))
	})
	return m
}

func useAuth(a *service.AuthService, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := utils.AuthMiddleware(r)
		if err != nil {
			utils.SendErrorResponse(w, 401, "unauthorized access")
			return
		}
		acc, err := a.UserStatus(info.UserID)
		if err != nil {
			utils.SendErrorResponse(w, 401, "unauthorized access")
			return
		}
		ctx := context.WithValue(r.Context(), "data", acc)
		f(w, r.WithContext(ctx))
	}
}
