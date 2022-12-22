package clint

import (
	"context"
	"log"

	"github.com/alramdein/user-service/config"
	"google.golang.org/grpc"

	"github.com/alramdein/user-service/pb"
)

func FindUserByUsernameAndPassword(ctx context.Context, username string, password string) error {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(config.GRPCPort(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
		return err
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	response, err := c.FindUserByUsernameAndPassword(ctx, &pb.FindUserByUsernameAndPasswordRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
		return err
	}

	log.Printf("Response from server: %s", response)
	return nil
}
