package user

import (
	"context"

	"github.com/kirillmc/trainings-auth/internal/model"
)

func (s *serv) LockUser(ctx context.Context, req *model.UserToLock) error {
	err := s.userRepository.LockUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
