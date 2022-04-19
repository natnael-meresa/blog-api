package persistence

import (
	"time"
	"twof/blog-api/internal/constant/model"

	"gorm.io/gorm"
)

type SubscriptionPersistence interface {
	CreateSubscription(subscription *model.Subscription) (err error)
	GetSubscriptionsByUser(subscription_date time.Time, userId uint) ([]model.Subscription, error)
}

type subscriptionPersistence struct {
	db *gorm.DB
}

func SubscriptionInit(db *gorm.DB) SubscriptionPersistence {
	return &subscriptionPersistence{
		db,
	}
}

func (sub *subscriptionPersistence) CreateSubscription(subscription *model.Subscription) (err error) {

	if err = sub.db.Create(subscription).Error; err != nil {
		return err
	}

	return nil
}

func (sub *subscriptionPersistence) GetSubscriptionsByUser(subscription_date time.Time, userId uint) ([]model.Subscription, error) {
	var subscriptions []model.Subscription
	if err := sub.db.Where("user_id = ? AND subscription_date >= ? ", userId, subscription_date).First(subscriptions).Error; err != nil {
		return nil, err
	}

	return subscriptions, nil
}

// func (ar *articlePersistence) GetArticleById(articleId uint, article *model.Article) (err error) {
// 	if err = ar.db.Where("ID = ?", articleId).First(article).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (ar *articlePersistence) GetAllArticles(article *[]model.Article) (err error) {
// 	if err = ar.db.Find(article).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
