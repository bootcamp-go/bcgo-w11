package internal

// SaleAttributes is the struct that represents the attributes of a sale.
type SaleAttributes struct {
	// Quantity is the quantity of the sale.
	Quantity int
	// ProductId is the product id of the sale.
	ProductId int
	// InvoiceId is the invoice id of the sale.
	InvoiceId int
}

// Sale is the struct that represents a sale.
type Sale struct {
	// Id is the unique identifier of the sale.
	Id int
	// SaleAttributes is the attributes of the sale.
	SaleAttributes
}