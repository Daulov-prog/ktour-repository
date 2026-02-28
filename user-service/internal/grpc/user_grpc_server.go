package grpc

import (
	"context"
	"user-service/internal/service"
	pb "user-service/proto/user"
)

type UserGRPCServer struct {
	pb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserGRPCServer(service *service.UserService) *UserGRPCServer {
	return &UserGRPCServer{service: service}
}

func (s *UserGRPCServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.service.GetUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		Id:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}
