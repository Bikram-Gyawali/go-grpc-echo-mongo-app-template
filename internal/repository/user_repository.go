package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type UserRepository struct {
	Collection *mongo.Collection
}



func (r *UserRepository) CreateUser(ctx context.Context, name, email, password string) (string, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindUserById(ctx context.Context, userID string) (*User, error) {
	var user User
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	err = r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
