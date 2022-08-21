package dto

import "time"

type GetOrderDetailsRequest struct {
	OrderID  string `json:"orderId" schema:"orderId"`
	SellerID int64  `json:"sellerId" schema:"sellerId"`
}

type GetOrderDetailsResponse struct {
	Code    int64                       `json:"code"`
	Message string                      `json:"message"`
	Data    GetOrderDetailsResponseData `json:"data"`
}

type GetOrderDetailsResponseData struct {
	TerminalCode          string                             `json:"terminalCode"`
	ReferralCode          string                             `json:"referralCode"`
	Status                string                             `json:"status"`
	OrderCaptureLineItems []OrderDetailsOrderCaptureLineItem `json:"orderCaptureLineItems"`
	Services              []OrderDetailsService              `json:"services"`
	OrderPromotion        []OrderDetailsOrderPromotion       `json:"orderPromotion"`
	TotalDiscount         int                                `json:"totalDiscount"`
	QrPromotionAmount     int                                `json:"qrPromotionAmount"`
	Payments              []OrderDetailsPayment              `json:"payments"`
	Refunds               []OrderDetailsRefund               `json:"refunds"`
	BillingInfo           OrderDetailsBillingInfo            `json:"billingInfo"`
	Shipments             []OrderDetailsShipment             `json:"shipments"`
	Returns               []OrderDetailsReturnRequest        `json:"returns"`
	Note                  string                             `json:"note"`
	Source                string                             `json:"source"`
	LoyaltyGrantPoint     int                                `json:"loyaltyGrantPoint"`
	LoyaltyGrantTierPoint int                                `json:"loyaltyGrantTierPoint"`
	OrderID               string                             `json:"orderId"`
	ReplacedOrderID       int64                              `json:"replacedOrderId"`
	TerminalID            int                                `json:"terminalId"`
	TerminalName          string                             `json:"terminalName"`
	PlatformID            int64                              `json:"platformId"`
	PlatformName          string                             `json:"platformName"`
	CreatedAt             int                                `json:"createdAt"`
	HoldStatus            bool                               `json:"holdStatus"`
	FulfillmentStatus     string                             `json:"fulfillmentStatus"`
	PaymentStatus         string                             `json:"paymentStatus"`
	GrandTotal            int                                `json:"grandTotal"`
	TotalPaid             int                                `json:"totalPaid"`
	RemainPayment         int                                `json:"remainPayment"`
	Creator               OrderDetailsCreator                `json:"creator"`
	Consultant            OrderDetailsConsultant             `json:"consultant"`
	Customer              OrderDetailsCustomer               `json:"customer"`
	ShippingInfo          OrderDetailsShippingInfo           `json:"shippingInfo"`
	ExternalOrderRef      string                             `json:"externalOrderRef"`
}

type OrderDetailsOrderCaptureLineItem struct {
	LineItemID        string                                      `json:"lineItemId"`
	Sku               string                                      `json:"sku"`
	SellerSku         string                                      `json:"sellerSku"`
	ServiceID         int64                                       `json:"serviceId"`
	DisplayName       string                                      `json:"displayName"`
	Uom               string                                      `json:"uom"`
	Quantity          int                                         `json:"quantity"`
	CancelledQuantity int                                         `json:"cancelledQuantity"`
	SiteID            int                                         `json:"siteId"`
	Serials           []string                                    `json:"serials"`
	Warranty          int                                         `json:"warranty"`
	Price             int                                         `json:"price"`
	OriginalPrice     int                                         `json:"originalPrice"`
	RowTotal          int                                         `json:"rowTotal"`
	IsGift            bool                                        `json:"isGift"`
	OnDemandDiscount  int                                         `json:"onDemandDiscount"`
	Promotions        []OrderDetailsOrderCaptureLineItemPromotion `json:"promotions"`
	IsAdult           bool                                        `json:"isAdult"`
}

type OrderDetailsOrderCaptureLineItemPromotion struct {
	ID               int                                                     `json:"id"`
	PromotionID      string                                                  `json:"promotionId"`
	Type             string                                                  `json:"type"`
	ApplyType        string                                                  `json:"applyType"`
	ApplyOn          []OrderDetailsOrderCaptureLineItemPromotionApplyOn      `json:"applyOn"`
	Discount         int                                                     `json:"discount"`
	OriginalDiscount int                                                     `json:"originalDiscount"`
	DiscountItems    []OrderDetailsOrderCaptureLineItemPromotionDiscountItem `json:"discountItems"`
	Gifts            []OrderDetailsOrderCaptureLineItemPromotionGift         `json:"gifts"`
	RemovedGifts     []OrderDetailsOrderCaptureLineItemPromotionRemovedGift  `json:"removedGifts"`
	Quantity         int                                                     `json:"quantity"`
	Voucher          OrderDetailsOrderCaptureLineItemPromotionVoucher        `json:"voucher"`
}

type OrderDetailsOrderCaptureLineItemPromotionApplyOn struct {
	LineItemID       string `json:"lineItemId"`
	ServiceID        int64  `json:"serviceId"`
	Quantity         int    `json:"quantity"`
	ComboConditionID string `json:"comboConditionId"`
	Sku              string `json:"sku"`
}

type OrderDetailsOrderCaptureLineItemPromotionDiscountItem struct {
	LineItemID string `json:"lineItemId"`
	Sku        string `json:"sku"`
	Quantity   int64  `json:"quantity"`
}

type OrderDetailsOrderCaptureLineItemPromotionGift struct {
	LineItemID string `json:"lineItemId"`
	Sku        string `json:"sku"`
	Quantity   int    `json:"quantity"`
	Name       string `json:"name"`
}

type OrderDetailsOrderCaptureLineItemPromotionRemovedGift struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type OrderDetailsOrderCaptureLineItemPromotionVoucher struct {
	Quantity     int    `json:"quantity"`
	Coupon       string `json:"coupon"`
	SellerIds    []int  `json:"sellerIds"`
	ComboApplyOn string `json:"comboApplyOn"`
}

type OrderDetailsService struct {
	ID         int64  `json:"id"`
	Sku        string `json:"sku"`
	Name       string `json:"name"`
	Fee        int64  `json:"fee"`
	SellerID   int    `json:"sellerId"`
	SellerName string `json:"sellerName"`
}

type OrderDetailsOrderPromotion struct {
	PromotionID int    `json:"promotionId"`
	Benefit     string `json:"benefit"`
}

type OrderDetailsPayment struct {
	PartnerCode            string `json:"partnerCode"`
	MethodCode             string `json:"methodCode"`
	OrderCode              string `json:"orderCode"`
	TransactionCode        string `json:"transactionCode"`
	ClientTransactionCode  string `json:"clientTransactionCode"`
	Amount                 string `json:"amount"`
	Status                 string `json:"status"`
	Message                string `json:"message"`
	MerchantMethodCode     string `json:"merchantMethodCode"`
	PartnerTransactionCode string `json:"partnerTransactionCode"`
	BankCode               string `json:"bankCode"`
	BankAccountNo          string `json:"bankAccountNo"`
	PaymentTerm            string `json:"paymentTerm"`
	CashierID              string `json:"cashierId"`
	PaidAt                 string `json:"paidAt"`
}

type OrderDetailsRefund struct {
	RefundID               int       `json:"refundId"`
	SellerID               int64     `json:"sellerId"`
	RequestAmount          int64     `json:"requestAmount"`
	RefundAmount           int64     `json:"refundAmount"`
	RefundPaymentReference string    `json:"refundPaymentReference"`
	PaymentMethod          string    `json:"paymentMethod"`
	Recipient              string    `json:"recipient"`
	Status                 string    `json:"status"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
	ConfirmedBy            string    `json:"confirmedBy"`
	ConfirmedAt            time.Time `json:"confirmedAt"`
}

type OrderDetailsBillingInfo struct {
	SplitInvoice              bool      `json:"splitInvoice"`
	SplitInvoiceGrandTotalMax int       `json:"splitInvoiceGrandTotalMax"`
	NewInvoiceDate            string    `json:"newInvoiceDate"`
	Name                      string    `json:"name"`
	Address                   string    `json:"address"`
	Email                     string    `json:"email"`
	TaxCode                   int64     `json:"taxCode"`
	Phone                     int       `json:"phone"`
	Type                      int       `json:"type"`
	Note                      string    `json:"note"`
	ExpectedIssueDate         time.Time `json:"expectedIssueDate"`
}

type OrderDetailsShipment struct {
	ShipmentID                  string                     `json:"shipmentId"`
	ShipmentCode                string                     `json:"shipmentCode"`
	SellerID                    int64                      `json:"sellerId"`
	Status                      string                     `json:"status"`
	Items                       []OrderDetailsShipmentItem `json:"items"`
	DocumentIds                 []string                   `json:"documentIds"`
	FulfillmentChannelID        int64                      `json:"fulfillmentChannelId"`
	FulfillmentChannelName      string                     `json:"fulfillmentChannelName"`
	IsAuto                      bool                       `json:"isAuto"`
	ReplacedShipmentID          string                     `json:"replacedShipmentId"`
	OutboundRequestID           string                     `json:"outboundRequestId"`
	OutboundRequestAt           int64                      `json:"outboundRequestAt"`
	WarehouseProviderExportID   string                     `json:"warehouseProviderExportId"`
	WarehouseExportAt           int64                      `json:"warehouseExportAt"`
	WarehouseProviderID         int64                      `json:"warehouseProviderId"`
	NumberOfPackages            int64                      `json:"numberOfPackages"`
	Height                      int64                      `json:"height"`
	Length                      int64                      `json:"length"`
	Width                       int64                      `json:"width"`
	TotalWeight                 int64                      `json:"totalWeight"`
	TransportProviderID         string                     `json:"transportProviderId"`
	TransportProviderName       string                     `json:"transportProviderName"`
	TransportProviderType       string                     `json:"transportProviderType"`
	TransportTrackingID         string                     `json:"transportTrackingId"`
	AllowTrialOnDelivery        string                     `json:"allowTrialOnDelivery"`
	DeliveryMethodCode          string                     `json:"deliveryMethodCode"`
	DeliveryRequestCreatedAt    time.Time                  `json:"deliveryRequestCreatedAt"`
	DeliveryRequestCreatorIamID string                     `json:"deliveryRequestCreatorIamId"`
	DeliveryRequestID           string                     `json:"deliveryRequestId"`
	DeliveryRequestStatus       string                     `json:"deliveryRequestStatus"`
	DeliveryTypeGroupID         string                     `json:"deliveryTypeGroupId"`
	DeliveryTypeID              string                     `json:"deliveryTypeId"`
	Description                 string                     `json:"description"`
	ExpectDeliveryDueFrom       time.Time                  `json:"expectDeliveryDueFrom"`
	ExpectDeliveryDueTo         time.Time                  `json:"expectDeliveryDueTo"`
}

type OrderDetailsShipmentItem struct {
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

type OrderDetailsReturnRequest struct {
	ReturnRequestID           int64                           `json:"returnRequestId"`
	PlatformID                string                          `json:"platformId"`
	ReasonID                  string                          `json:"reasonId"`
	Reason                    string                          `json:"reason"`
	SiteID                    int                             `json:"siteId"`
	WarehouseProviderExportID string                          `json:"warehouseProviderExportId"`
	Status                    string                          `json:"status"`
	Items                     []OrderDetailsReturnRequestItem `json:"items"`
	CustomerNotify            bool                            `json:"customerNotify"`
	CustomerInstruction       string                          `json:"customerInstruction"`
	Comment                   string                          `json:"comment"`
	CreatedByEmail            string                          `json:"createdByEmail"`
	ApprovedByEmail           string                          `json:"approvedByEmail"`
	CreatedAt                 time.Time                       `json:"createdAt"`
	UpdatedAt                 time.Time                       `json:"updatedAt"`
}

type OrderDetailsReturnRequestItem struct {
	ID                    string `json:"id"`
	SellerID              int64  `json:"sellerId"`
	ShipmentID            string `json:"shipmentId"`
	Sku                   string `json:"sku"`
	ServiceID             int    `json:"serviceId"`
	Name                  string `json:"name"`
	UnitPrice             int64  `json:"unitPrice"`
	Quantity              int64  `json:"quantity"`
	RequestReturnQuantity int64  `json:"requestReturnQuantity"`
	ReturnedQuantity      int64  `json:"returnedQuantity"`
	ReturnItemStatus      string `json:"returnItemStatus"`
}
