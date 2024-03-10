package utils

import (
	"ServerApp/service"
	"net/http"
)

type UseAuthFunc func(*service.AuthService, http.HandlerFunc) http.HandlerFunc
