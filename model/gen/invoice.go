package gen

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dokyan1989/postgresdemo/model"
)

func GenInvoice(orderID, shipmentID string, sellerID int64) *model.Invoice {
	now := time.Now()
	invoice := &model.Invoice{
		OrderID:    orderID,
		ShipmentID: shipmentID,
		SellerID:   sellerID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// gen invoice info
	var invoiceInfo model.InvoiceInfo
	gofakeit.Struct(&invoiceInfo)
	invoiceInfo.OrderID = orderID
	invoiceInfo.ShipmentID = shipmentID
	invoiceInfo.SellerID = sellerID

	invoice.InvoiceInfo = invoiceInfo

	return invoice
}
