package grpc

import (
	"github.com/alramdein/user-service/model"
	"github.com/alramdein/user-service/pb"
)

type Service struct {
	pb.UnimplementedUserServiceServer
	userUsecase model.UserUsecase
}

func NewUserService() *Service {
	return new(Service)
}

func (s *Service) RegisterUserUsecase(uc model.UserUsecase) {
	s.userUsecase = uc
}
