// internal/delivery/grpc/user_handler.go
package goclean

import (
	"context"
	"goclean/internal/domain"
	pb "goclean/proto"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userUseCase domain.UserUseCase
}

func NewUserHandler(userUseCase domain.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &domain.User{
		Name:    req.Name,
		Email:   req.Email,
		Age:     int(req.Age),
		Address: req.Address,
	}

	err := h.userUseCase.Create(user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUserResponse{
		User: toProtoUser(user),
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userUseCase.GetByID(uint(req.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.GetUserResponse{
		User: toProtoUser(user),
	}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := h.userUseCase.GetAll()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var protoUsers []*pb.User
	for _, user := range users {
		protoUsers = append(protoUsers, toProtoUser(&user))
	}

	return &pb.ListUsersResponse{
		Users: protoUsers,
	}, nil
}

func toProtoUser(user *domain.User) *pb.User {
	return &pb.User{
		Id:        uint64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Age:       int32(user.Age),
		Address:   user.Address,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}
