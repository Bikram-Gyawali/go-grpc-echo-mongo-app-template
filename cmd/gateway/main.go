package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "github.com/Bikram-Gyawali/grpc-echo-mongo-app/api/pb"
	"github.com/Bikram-Gyawali/grpc-echo-mongo-app/docs"
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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database(dbName)
	repos := repository.InitRepositories(db)
	services := services.InitServices(repos)

	go startGRPCServer(services)

	startRESTGateway()
}

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

	log.Println("gRPC server listening on", grpcAddr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func startRESTGateway() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover())

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost"+grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	gwmux := gw.NewServeMux()
	if err := pb.RegisterUserServiceHandler(context.Background(), gwmux, conn); err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	if err := docs.GenerateDocs(); err != nil {
		log.Fatalf("Failed to generate docs: %v", err)
	}

	e.Any("/*", echo.WrapHandler(gwmux))
	e.GET("/health", func(c echo.Context) error { return c.String(http.StatusOK, "Hello, World!") })
	e.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs/api")))))

	e.File("/swagger-ui", "docs/swagger-ui/index.html")
	e.GET("/swagger.json", func(c echo.Context) error { return c.File("docs/api/swagger.json") })

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowHeaders: []string{"*"},
	}))

	log.Println("REST API listening on", restAddr)
	e.Logger.Fatal(e.Start(restAddr))
}
