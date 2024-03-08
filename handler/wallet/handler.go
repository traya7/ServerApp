package wallet

import "github.com/gorilla/mux"

type WalletHandler struct{}

func New() *WalletHandler {
	return &WalletHandler{}
}

func (h *WalletHandler) Route(m *mux.Router) {
}
