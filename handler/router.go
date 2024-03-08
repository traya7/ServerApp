package handler

import (
	"AppServer/handler/auth"
	"AppServer/handler/wallet"
	"AppServer/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Params struct {
	Auth *service.AuthService
}

func NewRouter(p Params) http.Handler {
	m := mux.NewRouter()

	auth.New().Route(m)
	wallet.New().Route(m)

	return m
}
