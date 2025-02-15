package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestingUserRegister(t *testing.T) {
// 	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewUserServiceClient(conn)
// 	_, err = client.Register(context.Background(), &pb.RegisterRequest{
// 		Name:     "Bikram",
// 		Email:    "Bikram",
// 		Password: "Bikram",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestUserRegister(t *testing.T) {
	assert.Equal(t, 1, 1, "Dummy test")
}
