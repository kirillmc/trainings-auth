package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	descAccess "github.com/kirillmc/trainings-auth/pkg/access_v1"
)

var accessToken = flag.String("a", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM0NzI4MDgsImxvZ2luIjoiYSIsInJvbGUiOjN9.T6NSar3QKhsEFCJqntD5iUFSOa-BAC_AC8UhKZFLpto", "access token")

const (
	ExamplePath = "/user_v1.UserV1/Get"
)

const servicePort = 50055

func main() {
	flag.Parse()

	ctx := context.Background()
	md := metadata.New(map[string]string{"Authorization": "Bearer " + *accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%d", servicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial GRPC client: %v", err)
	}

	cl := descAccess.NewAccessV1Client(conn)

	_, err = cl.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: ExamplePath,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Access granted")
}
