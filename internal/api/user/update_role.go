package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kirillmc/trainings-auth/internal/converter"
	desc "github.com/kirillmc/trainings-auth/pkg/user_v1"
)

func (i *Implementation) UpdateRole(ctx context.Context, req *desc.UpdateRoleRequest) (*emptypb.Empty, error) {
	err := i.userService.UpdateRole(ctx, converter.ToRoleToUpdateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return nil, nil
}
