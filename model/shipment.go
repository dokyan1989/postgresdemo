package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Shipment struct {
	ID           string       `db:"id"`
	OrderID      string       `db:"order_id"`
	SellerID     int64        `db:"seller_id"`
	SiteID       int64        `db:"site_id"`
	Status       string       `db:"status"`
	ShipmentInfo ShipmentInfo `db:"shipment_info"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
}

type ShipmentInfo struct {
	OrderID                     string    `json:"orderId"`
	ShipmentID                  string    `json:"shipmentId"`
	ShipmentCode                string    `json:"shipmentCode"`
	SellerID                    int64     `json:"sellerId"`
	Status                      string    `json:"status"`
	Items                       []SIItem  `json:"items"`
	DocumentIds                 []string  `json:"documentIds"`
	FulfillmentChannelID        int64     `json:"fulfillmentChannelId"`
	FulfillmentChannelName      string    `json:"fulfillmentChannelName"`
	IsAuto                      bool      `json:"isAuto"`
	ReplacedShipmentID          string    `json:"replacedShipmentId"`
	OutboundRequestID           string    `json:"outboundRequestId"`
	OutboundRequestAt           int64     `json:"outboundRequestAt"`
	WarehouseProviderExportID   string    `json:"warehouseProviderExportId"`
	WarehouseExportAt           int64     `json:"warehouseExportAt"`
	WarehouseProviderID         int64     `json:"warehouseProviderId"`
	NumberOfPackages            int64     `json:"numberOfPackages"`
	Height                      int64     `json:"height"`
	Length                      int64     `json:"length"`
	Width                       int64     `json:"width"`
	TotalWeight                 int64     `json:"totalWeight"`
	TransportProviderID         string    `json:"transportProviderId"`
	TransportProviderName       string    `json:"transportProviderName"`
	TransportProviderType       string    `json:"transportProviderType"`
	TransportTrackingID         string    `json:"transportTrackingId"`
	AllowTrialOnDelivery        string    `json:"allowTrialOnDelivery"`
	DeliveryMethodCode          string    `json:"deliveryMethodCode"`
	DeliveryRequestCreatedAt    time.Time `json:"deliveryRequestCreatedAt"`
	DeliveryRequestCreatorIamID string    `json:"deliveryRequestCreatorIamId"`
	DeliveryRequestID           string    `json:"deliveryRequestId"`
	DeliveryRequestStatus       string    `json:"deliveryRequestStatus"`
	DeliveryTypeGroupID         string    `json:"deliveryTypeGroupId"`
	DeliveryTypeID              string    `json:"deliveryTypeId"`
	Description                 string    `json:"description"`
	ExpectDeliveryDueFrom       time.Time `json:"expectDeliveryDueFrom"`
	ExpectDeliveryDueTo         time.Time `json:"expectDeliveryDueTo"`
}

func (s ShipmentInfo) Value() (driver.Value, error) {
	valueString, err := json.Marshal(s)
	return string(valueString), err
}

func (s *ShipmentInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &s)
}

type SIItem struct {
	Sku          string   `json:"sku"`
	SiteID       int64    `json:"siteId"`
	LineItemID   string   `json:"lineItemId"`
	Name         string   `json:"name"`
	Services     []string `json:"services"`
	WarehouseID  string   `json:"warehouseId"`
	CancelledQty int64    `json:"cancelledQty"`
	DeliveredQty int64    `json:"deliveredQty"`
	ExportedQty  int64    `json:"exportedQty"`
	PackedQty    int64    `json:"packedQty"`
	RequestQty   int64    `json:"requestQty"`
	HoldID       string   `json:"holdId"`
	HoldItemID   string   `json:"holdItemId"`
	Priority     int64    `json:"priority"`
}
