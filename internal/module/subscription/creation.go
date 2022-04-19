package subscription

import (
	"fmt"
	"twof/blog-api/internal/constant/model"
)

func (s *service) CreateSubscription(subscription *model.Subscription) error {

	if err := s.subscriptionRepo.ValidateSubscription(subscription); err != nil {
		return err
	}

	err := s.subscriptionPersist.CreateSubscription(subscription)
	if err != nil {
		return fmt.Errorf("failed to save new Subscription %s", err.Error())
	}

	return nil
}
