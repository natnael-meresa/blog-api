package persistence

import (
	"twof/blog-api/internal/constant/model"

	"gorm.io/gorm"
)

type InvoicePersistence interface {
	CreateInvoice(invoice *model.Invoice) (err error)
}

type invoicePersistence struct {
	db *gorm.DB
}

func InvoiceInit(db *gorm.DB) InvoicePersistence {
	return &invoicePersistence{
		db,
	}
}

func (inv *invoicePersistence) CreateInvoice(invoice *model.Invoice) (err error) {

	if err = inv.db.Create(invoice).Error; err != nil {
		return err
	}

	return nil
}

// func (sub *subscriptionPersistence) GetSubscriptionsByUser(userId uint) ([]model.Subscription, error) {
// 	var subscriptions []model.Subscription
// 	if err := sub.db.Where("user_id = ?", userId).First(subscriptions).Error; err != nil {
// 		return nil, err
// 	}

// 	return subscriptions, nil
// }
