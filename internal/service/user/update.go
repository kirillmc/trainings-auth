package user

import (
	"context"
	"errors"

	"github.com/kirillmc/trainings-auth/internal/model"
)

func (s *serv) UpdateUser(ctx context.Context, req *model.UserToUpdate) error {
	err := s.userRepository.UpdateUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *serv) UpdatePassword(ctx context.Context, req *model.PasswordToUpdate) error {
	if req.ConfirmPassword != req.Password {
		return errors.New("Passwords are not equal")
	}

	err := s.userRepository.UpdatePassword(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
func (s *serv) UpdateRole(ctx context.Context, req *model.RoleToUpdate) error {
	err := s.userRepository.UpdateRole(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
