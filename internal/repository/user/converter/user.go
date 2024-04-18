package converter

import (
	"github.com/kirillmc/trainings-auth/internal/model"
	modelRepo "github.com/kirillmc/trainings-auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		Id:       user.Id,
		Login:    user.Login,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Name:     user.Name,
		Surname:  user.Surname,
		Role:     user.Role,
		IsLocked: user.IsLocked,
	}
}
