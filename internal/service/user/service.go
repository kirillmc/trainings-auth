package user

import (
	"github.com/kirillmc/trainings-auth/internal/repository"
	def "github.com/kirillmc/trainings-auth/internal/service"
)

var _ def.UserService = (*serv)(nil) //валидация имплементации интерфейса

type serv struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *serv {
	return &serv{userRepository: userRepository}
}
