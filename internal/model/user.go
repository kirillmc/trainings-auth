package model

import "google.golang.org/protobuf/types/known/wrapperspb"

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
	Login   *wrapperspb.StringValue
	Email   *wrapperspb.StringValue
	Name    *wrapperspb.StringValue
	Surname *wrapperspb.StringValue
	Avatar  *wrapperspb.StringValue
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
	Password        *wrapperspb.StringValue
	ConfirmPassword *wrapperspb.StringValue
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

const (
	RoleUnknown Role = 0
	RoleUser    Role = 1
	RoleModer   Role = 2
	RoleAdmin   Role = 3
)
