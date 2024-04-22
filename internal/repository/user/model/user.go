package model

import "github.com/kirillmc/trainings-auth/internal/model"

type User struct {
	Id       int64      `db:"id"`
	Name     string     `db:"name"`
	Surname  string     `db:"surname"`
	Email    string     `db:"email"`
	Avatar   string     `db:"avatar"`
	Login    string     `db:"login"`
	IsLocked bool       `db:"locked"`
	Role     model.Role `db:"role"`
	Weight   float64    `db:"weight"`
	Height   float64    `db:"height"`
}
