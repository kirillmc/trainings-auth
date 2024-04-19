package interceptor

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/kirillmc/trainings-auth/internal/service"
)

type Interceptor struct {
	Client service.AccessService
}

const (
	loginEndPoint      = "/auth_v1.AuthV1/Login"
	getAccessEndPoint  = "/auth_v1.AuthV1/GetAccessToken"
	getRefreshEndPoint = "/auth_v1.AuthV1/GetRefreshToken"
	checkEndPoint      = "/access_v1.AccessV1/Check"
	createEndPoint     = "/user_v1.UserV1/Create"
)

func (i *Interceptor) PolicyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}
	if info.FullMethod != loginEndPoint &&
		info.FullMethod != getAccessEndPoint &&
		info.FullMethod != getRefreshEndPoint &&
		info.FullMethod != checkEndPoint &&
		info.FullMethod != createEndPoint {
		err := i.Client.Check(metadata.NewOutgoingContext(ctx, md), info.FullMethod)
		if err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
