package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kirillmc/trainings-auth/internal/config/env"
	"github.com/kirillmc/trainings-auth/internal/model"
	"github.com/kirillmc/trainings-auth/internal/utils"
)

func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	refreshConfig, err := env.NewRefreshTokenConfig()
	if err != nil {
		return "", err
	}

	accessConfig, err := env.NewAccessTokenConfig()
	if err != nil {
		return "", err
	}

	claims, err := utils.VerifyToken(refreshToken, []byte(refreshConfig.RefreshTokenSecretKey()))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	accessToken, err := utils.GenerateToken(model.UserForToken{
		Login: claims.Login,
		Role:  claims.Role,
	},
		[]byte(accessConfig.AccessTokenSecretKey()),
		accessConfig.AccessTokenExpiration(),
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
