package auth

import (
	"github.com/kirillmc/trainings-auth/internal/repository"
	def "github.com/kirillmc/trainings-auth/internal/service"
)

var _ def.AuthService = (*serv)(nil) //валидация имплементации интерфейса

type serv struct {
	authRepository repository.AuthRepository
}

func NewService(authRepository repository.AuthRepository) *serv {
	return &serv{authRepository: authRepository}
}
