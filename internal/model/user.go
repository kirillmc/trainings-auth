package model

import "github.com/kirillmc/platform_common/pkg/nillable"

type User struct {
	Id       int64
	Login    string
	Email    string
	Role     Role
	Name     string
	Surname  string
	Avatar   string
	IsLocked bool
	Weight   float64
	Height   float64
}

type UserToCreate struct {
	Login           string
	Email           string
	Password        string
	ConfirmPassword string
	Role            Role
	Name            string
	Surname         string
	Avatar          string
}

type UserToUpdate struct {
	Id      int64
	Login   nillable.NilString
	Email   nillable.NilString
	Name    nillable.NilString
	Surname nillable.NilString
	Avatar  nillable.NilString
	Role    Role
}

type UserToLock struct {
	UserToLockId int64
}

type UserToUnlock struct {
	UserToUnlockId int64
}

type RoleToUpdate struct {
	UserId int64
	Role   Role
}

type PasswordToUpdate struct {
	UserId          int64
	Password        nillable.NilString
	ConfirmPassword nillable.NilString
}

type Anthropometry struct {
	UserId int64
	Weight nillable.NilDouble
	Height nillable.NilDouble
}

type UserToLogin struct {
	Login    string
	Password string
}

type UserForToken struct {
	Login string `json:"login"`
	Role  Role   `json:"role"`
}

type Role int32

func (r *Role) ToInt() int64 {
	switch *r {
	case RoleUser:
		return 1
	case RoleModer:
		return 2
	case RoleAdmin:
		return 3
	}

	return 0
}

const (
	RoleUnknown Role = 0
	RoleUser    Role = 1
	RoleModer   Role = 2
	RoleAdmin   Role = 3
)
