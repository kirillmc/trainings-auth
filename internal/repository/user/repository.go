package user

import (
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-auth/internal/repository"
)

// ТУТ ИМПЛЕМЕНТАЦИЯ МЕТОДОВ

const (
	usersTableName       = "users"
	accessRolesTableName = "roles_to_endpoints"
	adminsTableName      = "admins"
	modersTableName      = "moders"

	idColumn     = "id"
	userIdColumn = "user_id"

	nameColumn         = "name"
	surnameColumn      = "surname"
	loginColumn        = "login"
	emailColumn        = "email"
	avatarlColumn      = "avatar"
	passwordHashColumn = "password_hash"
	roleColumn         = "role"
	endpointColumn     = "endpoint"
	lockedColumn       = "locked"
	returnId           = "RETURNING id"
)

type repo struct {
	db db.Client
}

func NewUserRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}
func NewAuthRepository(db db.Client) repository.AuthRepository {
	return &repo{db: db}
}

func NewAccessRepository(db db.Client) repository.AccessRepository {
	return &repo{db: db}
}
