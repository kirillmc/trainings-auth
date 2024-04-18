package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kirillmc/trainings-auth/internal/converter"
	desc "github.com/kirillmc/trainings-auth/pkg/user_v1"
)

func (i *Implementation) LockUser(ctx context.Context, req *desc.LockUserRequest) (*emptypb.Empty, error) {
	err := i.userService.LockUser(ctx, converter.ToUserToLockFromDesc(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
