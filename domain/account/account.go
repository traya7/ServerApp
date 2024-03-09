package account

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID        string  `bson:"_id"`
	Username  string  `bson:"username"`
	Password  string  `bson:"password"`
	Balance   float64 `bson:"balance"`
	Role      string  `bson:"role"`
	CreatedBy string  `bson:"createdBy"`
	IsActive  bool    `bson:"isActive"`
}

func NewAccount(username, password, role, created_by string) Account {
	return Account{
		ID:        primitive.NewObjectID().Hex(),
		Username:  username,
		Password:  password,
		Balance:   0,
		Role:      role,
		CreatedBy: created_by,
		IsActive:  true,
	}
}
