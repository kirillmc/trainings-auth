package user

import (
	"context"

	//	"google.golang.org/grpc/metadata"

	"github.com/kirillmc/trainings-auth/internal/model"
	//"github.com/kirillmc/trainings-auth/internal/service/access/"
	//	descAccess "github.com/kirillmc/trainings-auth/pkg/access_v1"
	"github.com/pkg/errors"
)

func (s *serv) Create(ctx context.Context, req *model.UserToCreate) (int64, error) {
	//md, ok := metadata.FromIncomingContext(ctx)
	//if !ok {
	//	return 0, errors.New("metadata is not provided")
	//}
	//descAccess.AccessV1Client.Check(ctx, req)
	//err := access.Check(metadata.NewOutgoingContext(ctx, md), info.FullMethod)
	//if err != nil {
	//	return nil, err
	//}

	if req.ConfirmPassword != req.Password {
		return 0, errors.New("Passwords are not equal")
	}

	id, err := s.userRepository.Create(ctx, req)
	if err != nil {
		return 0, err
	}

	return id, nil
}
