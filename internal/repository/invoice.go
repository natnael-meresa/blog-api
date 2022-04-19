package repository

type InvoiceRepository interface {
	Generate() (err error)
}

type invoiceRepository struct {
}

func InvoiceInit() InvoiceRepository {
	return &invoiceRepository{}
}

func (sub *invoiceRepository) Generate() (err error) {
	return nil
}
