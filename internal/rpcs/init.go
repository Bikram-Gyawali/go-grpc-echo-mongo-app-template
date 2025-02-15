package rpcs

import (
	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/api/pb"
	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/internal/services"
	"google.golang.org/grpc"
)

type RPCServices struct {
	UserService *UserServiceServer
	BookService *BookServiceServer
}

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	UserService *services.UserService
}

type BookServiceServer struct {
	pb.UnimplementedBookServiceServer
	BookService *services.BookService
}

func InitRPCServices(s *grpc.Server, services *services.Services) (*grpc.Server, *RPCServices, error) {
	rpcServices := &RPCServices{
		UserService: &UserServiceServer{
			UserService: services.User,
		},
		// BookService: &BookServiceServer{
		// 	BookService: services.Book,
		// },
	}

	pb.RegisterUserServiceServer(s, rpcServices.UserService)
	// pb.RegisterBookServiceServer(s, rpcServices.BookService)

	return s, rpcServices, nil
}
