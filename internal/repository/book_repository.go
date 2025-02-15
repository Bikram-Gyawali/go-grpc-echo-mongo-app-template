package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

type BookRepository struct {
	Collection *mongo.Collection
}

func (r *BookRepository) CreateBook(ctx context.Context, name string) (string, error) {
	book := Book{
		Name: name,
	}
	result, err := r.Collection.InsertOne(ctx, book)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *BookRepository) FindBookById(ctx context.Context, bookID string) (*Book, error) {
	var book Book
	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, err
	}
	err = r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) FindBookByName(ctx context.Context, bookName string) (*Book, error) {
	var book Book
	err := r.Collection.FindOne(ctx, bson.M{"name": bookName}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) FindAllBooks(ctx context.Context) ([]Book, error) {
	var books []Book
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}
