package invoice

import (
	"fmt"
	"time"
	"twof/blog-api/internal/constant/model"

	"github.com/robfig/cron/v3"
)

func (s *service) Generate() error {

	c := cron.New()

	c.AddFunc("0 0 1 * *", func() {
		var users []model.User
		s.userPersist.GetAllUsers(&users)

		for _, user := range users {
			now := time.Now()
			subscription_date := now.AddDate(0, -1, 0)
			subscriptions, err := s.subscriptionPersist.GetSubscriptionsByUser(subscription_date, user.ID)

			if err != nil {
				fmt.Println("error fetching subscription")
			}
			var price = 0
			for range subscriptions {
				price += 200
			}

			invoice := model.Invoice{
				Price:       price,
				UserID:      user.ID,
				Description: "montly billing",
			}

			s.invoicePersistence.CreateInvoice(&invoice)
		}
	})

	c.Start()

	c.Stop()

	return nil
}
