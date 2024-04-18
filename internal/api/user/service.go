package user

import (
	"github.com/kirillmc/trainings-auth/internal/service"
	desc "github.com/kirillmc/trainings-auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
