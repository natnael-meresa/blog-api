package repository

import (
	"fmt"
	"twof/blog-api/internal/constant/model"

	"github.com/go-playground/validator/v10"
)

type SubscriptionRepository interface {
	ValidateSubscription(*model.Subscription) error
}

type subscriptionRepository struct {
}

func SubscriptionInit() SubscriptionRepository {
	return &subscriptionRepository{}
}

func (sub *subscriptionRepository) ValidateSubscription(subscription *model.Subscription) (err error) {
	validate = validator.New()

	err = validate.Struct(subscription)

	if err != nil {
		fmt.Println(err)
		var Msg string
		for _, err := range err.(validator.ValidationErrors) {
			Msg += err.Field() + " is " + err.Tag() + "\n"
		}
		return fmt.Errorf(Msg)
	}

	return nil
}
