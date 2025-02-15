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

const (
	mongoURI = "mongodb://mongo:27017"
	dbName   = "grpc-echo-mongo-app"
	grpcAddr = ":50051"
	restAddr = ":8080"
)

func main() {
	// Initialize MongoDB connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Initialize repositories and services
	db := client.Database(dbName)
	repos := repository.InitRepositories(db)
	services := services.InitServices(repos)

	// Start the gRPC server in a goroutine
	go startGRPCServer(services)

	// Start the REST API Gateway
	startRESTGateway()
}

// startGRPCServer initializes and runs the gRPC server
func startGRPCServer(services *services.Services) {
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	_, _, err = rpcs.InitRPCServices(s, services)
	if err != nil {
		log.Fatalf("Failed to register gRPC services: %v", err)
	}
	// pb.RegisterUserServiceServer(s, &rpcs.UserServiceServer{UserService: services.User})

	log.Println("gRPC server listening on", grpcAddr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// startRESTGateway initializes and runs the REST API using Echo + gRPC Gateway
func startRESTGateway() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	// Connect to gRPC Server
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost"+grpcAddr, // Ensure correct gRPC connection
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create gRPC Gateway
	gwmux := gw.NewServeMux()
	if err := pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	// Register routes with Echo
	e.Any("/*", echo.WrapHandler(gwmux))
	e.GET("/health", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, World!") })
	e.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs/api")))))

	log.Println("REST API listening on", restAddr)
	e.Logger.Fatal(e.Start(restAddr))
}
