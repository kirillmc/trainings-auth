package user

import (
	"context"

	"github.com/kirillmc/trainings-auth/internal/model"
)

func (s *serv) UnlockUser(ctx context.Context, req *model.UserToUnlock) error {
	err := s.userRepository.UnlockUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
