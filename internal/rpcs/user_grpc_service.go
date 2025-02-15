package rpcs

import (
	"context"
	"fmt"

	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/api/pb"
	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/middleware"
)

func (s *UserServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("Reached in gRPC response inside")
	userId, err := s.UserService.CreateUser(ctx, req.Name, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	fmt.Println("Reached in gRPC response agadi")
	return &pb.RegisterResponse{
		UserId:  userId,
		Message: "User registered successfully",
	}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userID, err := s.UserService.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := middleware.GenerateJWT(userID)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Message: "User logged in successfully",
		Token:   token,
	}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.UserService.FindUserById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
