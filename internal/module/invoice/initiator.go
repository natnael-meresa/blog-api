package invoice

import (
	"twof/blog-api/internal/repository"
	"twof/blog-api/internal/storage/persistence"
)

type Usecase interface {
	Generate() error
}

type service struct {
	invoiceRepo         repository.InvoiceRepository
	subscriptionPersist persistence.SubscriptionPersistence
	userPersist         persistence.UserPersistence
	invoicePersistence  persistence.InvoicePersistence
}

func Initialize(
	invoiceRepo repository.InvoiceRepository,
	subscriptionPersist persistence.SubscriptionPersistence,
	userPersist persistence.UserPersistence,
	invoicePersistence persistence.InvoicePersistence,
) Usecase {
	return &service{
		invoiceRepo,
		subscriptionPersist,
		userPersist,
		invoicePersistence,
	}
}
