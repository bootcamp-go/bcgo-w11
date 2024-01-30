package internal

// ServiceSale is the interface that wraps the basic ServiceSale methods.
type ServiceSale interface {
	// FindAll returns all sales.
	FindAll() (s []Sale, err error)
	// Save saves a sale.
	Save(s *Sale) (err error)
}