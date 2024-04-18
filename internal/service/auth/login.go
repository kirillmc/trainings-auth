package auth

import (
	"context"
	"errors"

	"github.com/kirillmc/platform_common/pkg/verify_password"
	"github.com/kirillmc/trainings-auth/internal/config/env"
	"github.com/kirillmc/trainings-auth/internal/model"
	"github.com/kirillmc/trainings-auth/internal/utils"
)

func (s *serv) Login(ctx context.Context, req *model.UserToLogin) (string, error) {
	// Сверяем хэши пароля
	hashPass, err := s.userRepository.GetHashPass(ctx, req.Login)
	if err != nil {
		return "", err
	}

	role, err := s.userRepository.GetRole(ctx, req.Login)
	if err != nil {
		return "", err
	}

	if !verify_password.VerifyPassword(hashPass, req.Password) {
		return "", errors.New("Wrong password!!!")
	}

	refreshConfig, err := env.NewRefreshTokenConfig()
	if err != nil {
		return "", err
	}

	refreshToken, err := utils.GenerateToken(
		model.UserForToken{
			Login: req.Login,
			Role:  role,
		},
		[]byte(refreshConfig.RefreshTokenSecretKey()),
		refreshConfig.RefreshTokenExpiration(),
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
