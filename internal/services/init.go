package services

import "github.com/Bikram-Gyawali/grpc-echo-mongo-app/internal/repository"

type Services struct {
	User *UserService
	Book *BookService
}

func InitServices(repos *repository.Repositories) *Services {
	return &Services{
		User: &UserService{
			UserRepo: repos.User,
		},
		Book: &BookService{
			BookRepo: repos.Book,
		},
	}
}
