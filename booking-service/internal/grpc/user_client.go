package grpc

import (
	pb "booking-service/proto/user"
	"context"

	"google.golang.org/grpc"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(addr string) (*UserClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &UserClient{client: pb.NewUserServiceClient(conn)}, nil
}

func (u *UserClient) GetUser(id string) (*pb.GetUserResponse, error) {
	return u.client.GetUser(context.Background(), &pb.GetUserRequest{Id: id})
}
