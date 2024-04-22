package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kirillmc/trainings-auth/internal/converter"
	desc "github.com/kirillmc/trainings-auth/pkg/user_v1"
)

func (i *Implementation) UpdateAnthropometry(ctx context.Context, req *desc.SetAnthropometryRequest) (*emptypb.Empty, error) {
	err := i.userService.UpdateAnthropometry(ctx, converter.ToAnthropometryFromDesc(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
