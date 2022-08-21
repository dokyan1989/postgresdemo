package gen

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dokyan1989/postgresdemo/model"
)

func GenInvoice(orderID string, shipmentID string, sellerID int64) *model.Invoice {
	invoice := &model.Invoice{
		OrderID:    orderID,
		ShipmentID: shipmentID,
		SellerID:   sellerID,
		CreatedAt:  gofakeit.DateRange(time.Now().Add((-3)*OneYear), time.Now()),
		UpdatedAt:  gofakeit.DateRange(time.Now().Add((-3)*OneYear), time.Now()),
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
