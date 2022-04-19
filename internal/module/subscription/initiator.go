package subscription

import (
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/repository"
	"twof/blog-api/internal/storage/persistence"
)

type Usecase interface {
	CreateSubscription(*model.Subscription) error
}

type service struct {
	subscriptionRepo    repository.SubscriptionRepository
	subscriptionPersist persistence.SubscriptionPersistence
}

func Initialize(
	subscriptionRepo repository.SubscriptionRepository,
	subscriptionPersist persistence.SubscriptionPersistence,
) Usecase {
	return &service{
		subscriptionRepo,
		subscriptionPersist,
	}
}
