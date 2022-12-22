package grpc

import (
	"context"

	"github.com/alramdein/user-service/pb"
	log "github.com/sirupsen/logrus"
)

func (s *Service) FindUserByUsernameAndPassword(ctx context.Context, req *pb.FindUserByUsernameAndPasswordRequest) (res *pb.User, err error) {
	user, err := s.userUsecase.FindByUsernameAndPassword(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		log.Error(err)
		return
	}

	res = user.ToProto()
	return
}
