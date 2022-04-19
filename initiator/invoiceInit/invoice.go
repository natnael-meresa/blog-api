package invoiceInit

import (
	"twof/blog-api/internal/module/invoice"
	"twof/blog-api/internal/repository"
	"twof/blog-api/internal/storage/persistence"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	databaseSubscription := persistence.SubscriptionInit(db)
	databaseUser := persistence.UserInit(db)
	databaseInvoice := persistence.InvoiceInit(db)

	repo := repository.InvoiceInit()
	usecase := invoice.Initialize(repo, databaseSubscription, databaseUser, databaseInvoice)

	go usecase.Generate()

}
