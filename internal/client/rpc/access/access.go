package access

import (
	"context"

	"github.com/kirillmc/trainings-auth/internal/client/rpc"
	descAccess "github.com/kirillmc/trainings-auth/pkg/access_v1"
)

type accessClient struct {
	client descAccess.AccessV1Client
}

var _ rpc.AccessClient = (*accessClient)(nil)

func NewAccessClient(client descAccess.AccessV1Client) rpc.AccessClient {
	return &accessClient{
		client: client,
	}
}

func (c *accessClient) Check(ctx context.Context, endpoint string) error {
	_, err := c.client.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: endpoint,
	})
	return err
}
