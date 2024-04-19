package model

type User struct {
	Id       int64
	Login    string
	Email    string
	Role     Role
	Name     string
	Surname  string
	Avatar   string
	IsLocked bool
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
	Login   NilString
	Email   NilString
	Name    NilString
	Surname NilString
	Avatar  NilString
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
	Password        NilString
	ConfirmPassword NilString
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

type NilString struct {
	Value   string
	IsEmpty bool
}

func Create(val string) NilString {
	if len(val) > 0 {
		return NilString{
			Value:   val,
			IsEmpty: false,
		}
	}
	return NilString{
		Value:   val,
		IsEmpty: true,
	}
}
