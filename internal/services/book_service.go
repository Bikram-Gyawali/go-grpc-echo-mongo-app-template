package services

import (
	"context"

	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/internal/repository"
)

type BookService struct {
	BookRepo *repository.BookRepository
}

func (s *BookService) CreateBook(ctx context.Context, name string) (string, error) {
	book := &repository.Book{
		Name: name,
	}
	result, err := s.BookRepo.CreateBook(ctx, book.Name)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *BookService) FindBookById(ctx context.Context, bookID string) (*repository.Book, error) {
	book, err := s.BookRepo.FindBookById(ctx, bookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) FindBookByName(ctx context.Context, bookName string) (*repository.Book, error) {
	book, err := s.BookRepo.FindBookByName(ctx, bookName)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *BookService) FindAllBooks(ctx context.Context) ([]repository.Book, error) {
	books, err := s.BookRepo.FindAllBooks(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
