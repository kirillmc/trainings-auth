package converter

import (
	"github.com/kirillmc/platform_common/pkg/nillable"
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
			Weight:   user.Weight,
			Height:   user.Height,
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
		Login:   nillable.Create(user.Info.Login.Value),
		Email:   nillable.Create(user.Info.Email.Value),
		Name:    nillable.Create(user.Info.Name.Value),
		Surname: nillable.Create(user.Info.Surname.Value),
		Avatar:  nillable.Create(user.Info.Avatar.Value),
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
		Password:        nillable.Create(password.Info.Password.Value),
		ConfirmPassword: nillable.Create(password.Info.PasswordConfirm.Value),
	}
}
func ToAnthropometryFromDesc(anthropometry *desc.SetAnthropometryRequest) *model.Anthropometry {
	return &model.Anthropometry{
		UserId: anthropometry.UserId,
		Weight: nillable.CreateNillableDouble(anthropometry.Weight.Value),
		Height: nillable.CreateNillableDouble(anthropometry.Height.Value),
	}
}
func ToUserToLoginFromDescAuth(user *descAuth.LoginRequest) *model.UserToLogin {
	return &model.UserToLogin{
		Login:    user.Login,
		Password: user.Password,
	}
}
