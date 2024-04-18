package converter

import (
	"github.com/kirillmc/trainings-auth/internal/model"
	descAuth "github.com/kirillmc/trainings-auth/pkg/auth_v1"
	desc "github.com/kirillmc/trainings-auth/pkg/user_v1"
)

func ToGetResponseFromService(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		Id: user.Id,
		User: &desc.User{
			Login:    user.Login,
			Email:    user.Email,
			Role:     desc.Role(user.Role),
			Name:     user.Name,
			Surname:  user.Surname,
			Avatar:   user.Avatar,
			IsLocked: user.IsLocked,
		},
	}
}

func ToUserModelCreateFromDesc(user *desc.CreateRequest) *model.UserToCreate {
	return &model.UserToCreate{
		Login:           user.User.Login,
		Email:           user.User.Email,
		Password:        user.User.Password,
		ConfirmPassword: user.User.PasswordConfirm,
		Role:            model.Role(user.User.Role),
		Name:            user.User.Name,
		Surname:         user.User.Surname,
		Avatar:          user.User.Avatar,
	}
}

func ToUserModelUpdateFromDesc(user *desc.UpdateRequest) *model.UserToUpdate {
	return &model.UserToUpdate{
		Id:      user.Id,
		Login:   user.Info.Login,
		Email:   user.Info.Email,
		Name:    user.Info.Name,
		Surname: user.Info.Surname,
		Avatar:  user.Info.Avatar,
	}
}

func ToRoleToUpdateFromDesc(role *desc.UpdateRoleRequest) *model.RoleToUpdate {
	return &model.RoleToUpdate{
		UserId: role.UserId,
		Role:   model.Role(role.Role),
	}
}

func ToUserToLockFromDesc(lock *desc.LockUserRequest) *model.UserToLock {
	return &model.UserToLock{
		UserToLockId: lock.UserToLockId,
	}
}
func ToUserToUnlockFromDesc(unlock *desc.UnlockUserRequest) *model.UserToUnlock {
	return &model.UserToUnlock{
		UserToUnlockId: unlock.UserToUnlockId,
	}
}

func ToPasswordToUpdateFromDesc(password *desc.UpdatePasswordRequest) *model.PasswordToUpdate {
	return &model.PasswordToUpdate{
		UserId:          password.UserId,
		Password:        password.Info.Password,
		ConfirmPassword: password.Info.PasswordConfirm,
	}
}
func ToUserToLoginFromDescAuth(user *descAuth.LoginRequest) *model.UserToLogin {
	return &model.UserToLogin{
		Login:    user.Login,
		Password: user.Password,
	}
}
