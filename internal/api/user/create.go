package user

import (
	"context"

	"github.com/kirillmc/trainings-auth/internal/converter"
	desc "github.com/kirillmc/trainings-auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserModelCreateFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
