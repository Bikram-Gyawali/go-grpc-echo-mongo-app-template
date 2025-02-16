package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, name, email, password string) (string, error) {
	existingUser, err := s.UserRepo.FindUserByEmail(ctx, email)
	if err != nil && err != mongo.ErrNoDocuments {
		return "", err 
	}
	fmt.Println("existingUser", existingUser)
	if existingUser != nil {
		fmt.Println("Email already exists")
		return "", errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return s.UserRepo.CreateUser(ctx, name, email, string(hashedPassword))
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (*repository.User, error) {
	return s.UserRepo.FindUserByEmail(ctx, email)
}

func (s *UserService) FindUserById(ctx context.Context, userID string) (*repository.User, error) {
	return s.UserRepo.FindUserById(ctx, userID)
}

func (s *UserService) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := s.UserRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
