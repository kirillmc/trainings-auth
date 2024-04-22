package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/trainings-auth/internal/model"
	"github.com/kirillmc/trainings-auth/internal/repository/user/converter"
	modelRepo "github.com/kirillmc/trainings-auth/internal/repository/user/model"
)

// ТУТ ИМПЛЕМЕНТАЦИЯ МЕТОДОВ

func (r *repo) Create(ctx context.Context, req *model.UserToCreate) (int64, error) {
	hashPass, err := genPassHash(req.Password)
	if err != nil {
		return 0, err
	}

	builder := sq.Insert(usersTableName).PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, surnameColumn, emailColumn, avatarlColumn, loginColumn, passwordHashColumn, roleColumn).
		Values(req.Name, req.Surname, req.Email, req.Avatar, req.Login, hashPass, req.Role).
		Suffix(returnId)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	// Добавлено с db.Client
	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) GetUser(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, surnameColumn, emailColumn, avatarlColumn, loginColumn, lockedColumn, roleColumn, weightColumn, heightColumn).
		PlaceholderFormat(sq.Dollar).
		From(usersTableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetUser",
		QueryRaw: query,
	}

	var user modelRepo.User

	err = r.db.DB().ScanOneContext(ctx, &user, q, args...) // Сканирует одну запись в user
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) UpdateUser(ctx context.Context, req *model.UserToUpdate) error {
	builder := sq.Update(usersTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: req.Id})

	if !req.Login.IsEmpty {
		builder = builder.Set(loginColumn, req.Login.Value)
	}

	if !req.Email.IsEmpty {
		builder = builder.Set(emailColumn, req.Email.Value)
	}

	if !req.Name.IsEmpty {
		builder = builder.Set(nameColumn, req.Name.Value)
	}
	if !req.Surname.IsEmpty {
		builder = builder.Set(surnameColumn, req.Surname.Value)
	}

	if !req.Avatar.IsEmpty {
		builder = builder.Set(avatarlColumn, req.Avatar.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UpdateUser",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(usersTableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdatePassword(ctx context.Context, req *model.PasswordToUpdate) error {
	builder := sq.Update(usersTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: req.UserId})

	if !req.Password.IsEmpty {
		hashPass, err := genPassHash(req.Password.Value)
		if err != nil {
			return err
		}

		builder = builder.Set(passwordHashColumn, hashPass)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UpdatePassword",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *repo) UpdateRole(ctx context.Context, req *model.RoleToUpdate) error {
	builder := sq.Update(usersTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: req.UserId})

	if req.Role != model.RoleUnknown {
		builder = builder.Set(roleColumn, req.Role)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UpdateRole",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) LockUser(ctx context.Context, req *model.UserToLock) error {
	builder := sq.Update(usersTableName).
		PlaceholderFormat(sq.Dollar).Set(lockedColumn, true).
		Where(sq.Eq{idColumn: req.UserToLockId})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.LockUser",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *repo) UnlockUser(ctx context.Context, req *model.UserToUnlock) error {
	builder := sq.Update(usersTableName).
		PlaceholderFormat(sq.Dollar).Set(lockedColumn, false).
		Where(sq.Eq{idColumn: req.UserToUnlockId})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UnlockUser",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateAnthropometry(ctx context.Context, req *model.Anthropometry) error {
	builder := sq.Update(usersTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: req.UserId})

	if !req.Height.IsEmpty {
		builder = builder.Set(heightColumn, req.Height.Value)
	}

	if !req.Weight.IsEmpty {
		builder = builder.Set(weightColumn, req.Weight.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.UpdateAnthropometry",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func genPassHash(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}
