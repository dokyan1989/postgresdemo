package gen

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dokyan1989/postgresdemo/helper/random"
	"github.com/dokyan1989/postgresdemo/model"
	"github.com/google/uuid"
)

type GenShipmentOption func(s *model.Shipment)

func GenShipmentWithStatus(status string) GenShipmentOption {
	return func(s *model.Shipment) {
		if status != "" {
			s.Status = status
		}
	}
}
func GenShipment(orderID string, opts ...GenShipmentOption) *model.Shipment {
	now := time.Now()
	shipment := &model.Shipment{
		ID:        uuid.NewString(),
		OrderID:   orderID,
		SellerID:  int64(random.Int(1, 10)),
		SiteID:    int64(random.Int(1, 10)),
		Status:    GenShipmentStatus(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	for _, opt := range opts {
		opt(shipment)
	}

	// gen shipment info
	var shipmentInfo model.ShipmentInfo
	gofakeit.Struct(&shipmentInfo)
	shipmentInfo.OrderID = orderID
	shipmentInfo.ShipmentID = shipment.ID
	shipmentInfo.SellerID = shipment.SellerID

	shipment.ShipmentInfo = shipmentInfo

	return shipment
}

func GenShipmentStatus() string {
	values := []string{
		"NEW",
		"WAITING_FOR_MANUAL_CONFIRMATION",
		"PROVIDER_TRANSFERRING",
		"PROVIDER_TRANSFERRED",
		"WAITING_FOR_PRODUCT",
		"PICKED",
		"PACKED",
		"DELIVERING",
		"DELIVERED",
		"PARTIAL_DELIVERED",
		"CANCELLED",
		"RETURNED",
		"PARTIAL_CANCELLED",
	}

	i := random.Int(0, len(values)-1)
	return values[i]
}
