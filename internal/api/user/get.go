package user

import (
	"context"

	"github.com/kirillmc/trainings-auth/internal/converter"
	desc "github.com/kirillmc/trainings-auth/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	nUser, err := i.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetResponseFromService(nUser), nil
}
