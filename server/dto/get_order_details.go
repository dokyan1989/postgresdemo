package dto

import "time"

type GetOrderDetailsRequest struct {
	OrderID  string `json:"orderId" schema:"orderId"`
	SellerID int64  `json:"sellerId" schema:"sellerId"`
}

type GetOrderDetailsResponse struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Data    *FullSaleOrder `json:"data"`
}

type FullSaleOrder struct {
	TerminalCode          string                 `json:"terminalCode"`
	ReferralCode          string                 `json:"referralCode"`
	Status                string                 `json:"status"`
	OrderCaptureLineItems []OrderCaptureLineItem `json:"orderCaptureLineItems"`
	Services              []OrderService         `json:"services"`
	OrderPromotion        []OrderPromotion       `json:"orderPromotion"`
	TotalDiscount         float64                `json:"totalDiscount"`
	QrPromotionAmount     float64                `json:"qrPromotionAmount"`
	Payments              []OrderPayment         `json:"payments"`
	Refunds               []OrderRefund          `json:"refunds"`
	BillingInfo           OrderBillingInfo       `json:"billingInfo"`
	Shipments             []Shipment             `json:"shipments"`
	Returns               []ReturnInfo           `json:"returns"`
	Note                  string                 `json:"note"`
	Source                string                 `json:"source"`
	LoyaltyGrantPoint     int                    `json:"loyaltyGrantPoint"`
	LoyaltyGrantTierPoint int                    `json:"loyaltyGrantTierPoint"`
	OrderID               string                 `json:"orderId"`
	ReplacedOrderID       string                 `json:"replacedOrderId"`
	TerminalID            int64                  `json:"terminalId"`
	TerminalName          string                 `json:"terminalName"`
	PlatformID            int64                  `json:"platformId"`
	PlatformName          string                 `json:"platformName"`
	CreatedAt             time.Time              `json:"createdAt"`
	HoldStatus            bool                   `json:"holdStatus"`
	FulfillmentStatus     string                 `json:"fulfillmentStatus"`
	PaymentStatus         string                 `json:"paymentStatus"`
	GrandTotal            float64                `json:"grandTotal"`
	TotalPaid             float64                `json:"totalPaid"`
	RemainPayment         float64                `json:"remainPayment"`
	Creator               OrderCreator           `json:"creator"`
	Consultant            OrderConsultant        `json:"consultant"`
	Customer              OrderCustomer          `json:"customer"`
	ShippingInfo          OrderShippingInfo      `json:"shippingInfo"`
	ExternalOrderRef      string                 `json:"externalOrderRef"`
}
