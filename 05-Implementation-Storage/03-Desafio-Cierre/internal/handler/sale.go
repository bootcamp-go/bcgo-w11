package handler

import (
	"net/http"

	"app/internal"
	"app/platform/web/request"
	"app/platform/web/response"
)

// NewSalesDefault returns a new SalesDefault
func NewSalesDefault(sv internal.ServiceSale) *SalesDefault {
	return &SalesDefault{sv: sv}
}

// SalesDefault is a struct that returns the sale handlers
type SalesDefault struct {
	// sv is the sale's service
	sv internal.ServiceSale
}

// SaleJSON is a struct that represents a sale in JSON format
type SaleJSON struct {
	Id int `json:"id"`
	Quantity int `json:"quantity"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
}

// GetAll returns all sales
func (h *SalesDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		s, err := h.sv.FindAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting sales")
			return
		}

		// response
		// - serialize
		sJSON := make([]SaleJSON, len(s))
		for ix, v := range s {
			sJSON[ix] = SaleJSON{
				Id:        v.Id,
				Quantity: v.Quantity,
				ProductId:  v.ProductId,
				InvoiceId: v.InvoiceId,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "sales found",
			"data":    sJSON,
		})
	}
}

// RequestBodySale is a struct that represents the request body for a sale
type RequestBodySale struct {
	Quantity int `json:"quantity"`
	ProductId int `json:"product_id"`
	InvoiceId int `json:"invoice_id"`
}
// Create creates a new sale
func (h *SalesDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodySale
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing request body")
			return
		}

		// process
		// - deserialize
		s := internal.Sale{
			SaleAttributes: internal.SaleAttributes{
				Quantity: reqBody.Quantity,
				ProductId: reqBody.ProductId,
				InvoiceId: reqBody.InvoiceId,
			},
		}
		// - save
		err = h.sv.Save(&s)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving sale")
			return
		}

		// response
		// - serialize
		sa := SaleJSON{
			Id:        s.Id,
			Quantity: s.Quantity,
			ProductId:  s.ProductId,
			InvoiceId: s.InvoiceId,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "sale created",
			"data":    sa,
		})
	}
}