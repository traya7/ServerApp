package service

import (
	"ServerApp/domain/account"
	"ServerApp/types"
)

type AuthService struct {
	repo account.Repository
}

func NewAuthService(r account.Repository) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (*AuthService) UserLogin(username, password, role string) (*types.User, error) {
	return nil, nil
}
func (*AuthService) UserStatus(user_id string) (*types.User, error) {
	return nil, nil
}
func (*AuthService) UserResetPwd() {}
