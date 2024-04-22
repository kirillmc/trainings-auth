package auth

import (
	"context"
	"errors"

	"github.com/kirillmc/platform_common/pkg/verify_password"
	"github.com/kirillmc/trainings-auth/internal/config/env"
	"github.com/kirillmc/trainings-auth/internal/model"
	"github.com/kirillmc/trainings-auth/internal/utils"
)

func (s *serv) Login(ctx context.Context, req *model.UserToLogin) (string, int64, error) {
	// Сверяем хэши пароля
	hashPass, err := s.authRepository.GetHashPass(ctx, req.Login)
	if err != nil {
		return "", 0, err
	}

	role, err := s.authRepository.GetRole(ctx, req.Login)
	if err != nil {
		return "", 0, err
	}

	if !verify_password.VerifyPassword(hashPass, req.Password) {
		return "", 0, errors.New("Wrong password!!!")
	}

	refreshConfig, err := env.NewRefreshTokenConfig()
	if err != nil {
		return "", 0, err
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
		return "", 0, errors.New("failed to generate token")
	}

	user_id, err := s.authRepository.GetUserIdByLoginAndPass(ctx, req.Login, hashPass)
	if err != nil {
		return "", 0, err
	}

	return refreshToken, user_id, nil
}
