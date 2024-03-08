package account

import "ServerApp/domain"

type mongoRepo struct{}

func NewMongoRepo(db *domain.MongoDB) *mongoRepo {
	return &mongoRepo{}
}

func (r *mongoRepo) GetAccountById(user_id string) (*Account, error) {
	return nil, nil
}

func (r *mongoRepo) GetAccountByUsername(username string) (*Account, error) {
	return nil, nil
}
