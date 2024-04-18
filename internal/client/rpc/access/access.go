package access

import (
	"context"

	"github.com/kirillmc/trainings-auth/internal/client/rpc"
	"github.com/kirillmc/trainings-auth/internal/service"
)

type accessClient struct {
	client service.AccessService
}

var _ rpc.AccessClient = (*accessClient)(nil)

func NewAccessClient(client service.AccessService) rpc.AccessClient {
	return &accessClient{
		client: client,
	}
}

func (c *accessClient) Check(ctx context.Context, endpoint string) error {
	err := c.client.Check(ctx, endpoint)

	return err
}
