package service

import (
	"context"

	"github.com/kirillmc/trainings-auth/internal/model"
)

type UserService interface {
	Create(ctx context.Context, req *model.UserToCreate) (int64, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, req *model.UserToUpdate) error
	UpdatePassword(ctx context.Context, req *model.PasswordToUpdate) error
	UpdateRole(ctx context.Context, req *model.RoleToUpdate) error
	LockUser(ctx context.Context, req *model.UserToLock) error
	UnlockUser(ctx context.Context, req *model.UserToUnlock) error
}

type AccessService interface {
	Check(ctx context.Context, endpointAddress string) error
}

type AuthService interface {
	Login(ctx context.Context, req *model.UserToLogin) (string, error)
	GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}
