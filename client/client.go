package clint

import (
	"context"
	"time"

	grpcpool "github.com/processout/grpc-go-pool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/alramdein/user-service/pb"
)

type userClient struct {
	Conn *grpcpool.Pool
}

func NewClient(target string, timeout time.Duration, idleConnPool, maxConnPool int) (pb.UserServiceClient, error) {
	factory := newFactory(target, timeout)

	pool, err := grpcpool.New(factory, idleConnPool, maxConnPool, time.Second)
	if err != nil {
		log.Errorf("Error : %v", err)
		return nil, err
	}

	return &userClient{
		Conn: pool,
	}, nil
}

func newFactory(target string, timeout time.Duration) grpcpool.Factory {
	return func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(target, grpc.WithInsecure(), withClientUnaryInterceptor(timeout))
		if err != nil {
			log.Errorf("Error : %v", err)
			return nil, err
		}

		return conn, err
	}
}

func withClientUnaryInterceptor(timeout time.Duration) grpc.DialOption {
	return grpc.WithUnaryInterceptor(func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	})
}

func (u *userClient) FindUserByUsernameAndPassword(ctx context.Context, in *pb.FindUserByUsernameAndPasswordRequest, opts ...grpc.CallOption) (*pb.User, error) {
	conn, err := u.Conn.Get(ctx)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
		return nil, err
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	user, err := c.FindUserByUsernameAndPassword(ctx, &pb.FindUserByUsernameAndPasswordRequest{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		log.Fatalf("Error when calling FindUserByUsernameAndPassword: %s", err)
		return nil, err
	}

	log.Printf("Response from server: %s", user)
	return user, nil
}
