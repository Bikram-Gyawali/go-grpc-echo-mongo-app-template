package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repositories struct {
	User *UserRepository
	Book *BookRepository
}

func InitRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		User: &UserRepository{
			Collection: db.Collection("users"),
		},
		Book: &BookRepository{
			Collection: db.Collection("books"),
		},
	}
}
