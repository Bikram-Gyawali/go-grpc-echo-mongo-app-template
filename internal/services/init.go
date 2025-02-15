package services

import "github.com/Bikram-Gyawali/grpc-echo-mongo-app/internal/repository"

type Services struct {
	User *UserService
	Book *BookService
}

func InitServices(db *repository.Repositories) *Services {
	return &Services{
		User: &UserService{
			UserRepo: db.User,
		},
		Book: &BookService{
			BookRepo: db.Book,
		},
	}
}
