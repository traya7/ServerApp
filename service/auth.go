package service

import (
	"ServerApp/domain/account"
	"ServerApp/types"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized access")
	ErrInternalError      = errors.New("internal error")
)

type AuthService struct {
	repo account.Repository
}

func NewAuthService(r account.Repository) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (s *AuthService) CreateAccount(username, password, admin_id, admin_role string) error {
	//

	role := subRole(admin_role)
	if role == "" {
		return ErrUnauthorized
	}

	// hashing password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ErrInternalError
	}
	hashedPassword := string(bytes)

	// save account in db
	acc := account.NewAccount(username, hashedPassword, role, admin_id)
	if err := s.repo.Store(acc); err != nil {
		return ErrInternalError
	}

	return nil
}

func (s *AuthService) UserLogin(username, password, role string) (*types.User, error) {

	acc, err := s.repo.FindOneByUsername(username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if role == "USER" && acc.Role != role {
		return nil, errors.New("account with no requested role.")
	}

	if !acc.IsActive {
		return nil, errors.New("account is blocked.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &types.User{
		ID:       acc.ID,
		Username: acc.Username,
		Balance:  acc.Balance,
		Role:     acc.Role,
		ImgUri:   acc.ImgUri,
	}, nil
}

func (s *AuthService) UserStatus(user_id string) (*types.User, error) {
	acc, err := s.repo.FindOneByID(user_id)
	if err != nil || !acc.IsActive {
		return nil, ErrUnauthorized
	}

	return &types.User{
		ID:       acc.ID,
		Username: acc.Username,
		Balance:  acc.Balance,
		Role:     acc.Role,
		ImgUri:   acc.ImgUri,
	}, nil
}

func (s *AuthService) UpdateMyPwd(user_id, old_pwd, new_pwd string) error {
	acc, err := s.repo.FindOneByID(user_id)
	if err != nil || !acc.IsActive {
		return ErrUnauthorized
	}
	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(old_pwd))
	if err != nil {
		return errors.New("invalid password")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(new_pwd), 14)
	if err != nil {
		return ErrInternalError
	}
	acc.Password = string(bytes)
	if err := s.repo.Update(*acc); err != nil {
		return err
	}
	return nil
}

func (*AuthService) UserResetPwd() {}

func subRole(r string) string {
	if r == "SUPER" {
		return "ADMIN"
	}
	if r == "ADMIN" {
		return "SHOP"
	}
	if r == "SHOP" {
		return "USER"
	}
	return ""
}
