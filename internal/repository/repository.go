package repository

import (
	"context"

	"github.com/kirillmc/trainings-auth/internal/model"
)

// файл ТОЛЬКО для интерфейсов

type UserRepository interface {
	Create(ctx context.Context, req *model.UserToCreate) (int64, error)

	GetUser(ctx context.Context, id int64) (*model.User, error)

	UpdateUser(ctx context.Context, req *model.UserToUpdate) error
	UpdatePassword(ctx context.Context, req *model.PasswordToUpdate) error
	UpdateRole(ctx context.Context, req *model.RoleToUpdate) error
	UpdateAnthropometry(ctx context.Context, req *model.Anthropometry) error

	LockUser(ctx context.Context, req *model.UserToLock) error
	UnlockUser(ctx context.Context, req *model.UserToUnlock) error
	Delete(ctx context.Context, id int64) error
}

type AuthRepository interface {
	GetRole(ctx context.Context, login string) (model.Role, error)
	GetHashPass(ctx context.Context, login string) (string, error)
	GetUserIdByLoginAndPass(ctx context.Context, login string, pass string) (int64, error)
}

type AccessRepository interface {
	GetAccessibleRoles(ctx context.Context) (map[string]model.Role, error)
}
