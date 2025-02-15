package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/Bikram-Gyawali/grpc-echo-mongo-app/api/pb"
	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/internal/repository"
	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/internal/rpcs"
	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	gw "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Start the gRPC server
	go startGRPCServer()

	// Create Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to gRPC Server
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create gRPC Gateway
	gwmux := gw.NewServeMux()
	err = pb.RegisterUserServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	// Register Gateway with Echo
	e.Any("/*", echo.WrapHandler(gwmux))

	e.GET("/health", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, World!") })

	e.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs/api")))))

	// Start HTTP server
	log.Println("REST API listening on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}

// startGRPCServer initializes and runs the gRPC server
func startGRPCServer() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("grpc-echo-mongo-app")

	userRepo := &repository.UserRepository{
		Collection: db.Collection("users"),
	}

	userService := &services.UserService{
		UserRepo: userRepo,
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &rpcs.UserServiceServer{
		UserService: userService,
	})

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
