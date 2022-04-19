package user

import (
	"fmt"
	"twof/blog-api/internal/constant/model"
)

func (s *service) ValidateUser(user *model.User) error {
	err := s.userRepo.ValidateUser(user)
	if err != nil {
		return fmt.Errorf("failed to Validate user %s", err.Error())
	}

	return nil
}
