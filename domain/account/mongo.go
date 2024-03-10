package account

import (
	"ServerApp/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	col *mongo.Collection
}

func NewMongoRepo(db *domain.MongoDB) *mongoRepo {
	return &mongoRepo{
		col: db.GetCollection("accounts"),
	}
}

func (r *mongoRepo) Store(acc Account) error {
	if _, err := r.col.InsertOne(context.TODO(), acc); err != nil {
		return err
	}
	return nil
}
func (r *mongoRepo) Update(acc Account) error {
	filter := bson.D{{Key: "_id", Value: acc.ID}}
	if _, err := r.col.UpdateOne(context.TODO(), filter, acc); err != nil {
		return err
	}
	return nil
}
func (r *mongoRepo) FindOneByID(user_id string) (*Account, error) {
	filter := bson.D{{Key: "_id", Value: user_id}}
	var result Account
	err := r.col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *mongoRepo) FindOneByUsername(username string) (*Account, error) {
	filter := bson.D{{Key: "username", Value: username}}
	var result Account
	err := r.col.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
