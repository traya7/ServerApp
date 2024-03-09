package handler

import (
	"ServerApp/handler/auth"
	"ServerApp/handler/utils"
	"ServerApp/handler/wallet"
	"ServerApp/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Params struct {
	Auth *service.AuthService
}

func NewRouter(p Params) http.Handler {
	m := mux.NewRouter()

	auth.New(p.Auth).Route(m.PathPrefix("/v1/auth").Subrouter())
	wallet.New().Route(m.PathPrefix("/v1/wallet").Subrouter())

	m.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.SendErrorResponse(w, 404, "Opps, not found page.")
	})
	return m
}
