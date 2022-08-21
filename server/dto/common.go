package dto

import "time"

type ListOrdersResponseData struct {
	Orders     []BaseSaleOrder `json:"orders"`
	NextOffset int             `json:"nextOffset"`
}

type BaseSaleOrder struct {
	OrderID           string            `json:"orderId"`
	ReplacedOrderID   string            `json:"replacedOrderId"`
	TerminalID        int64             `json:"terminalId"`
	TerminalName      string            `json:"terminalName"`
	PlatformID        int64             `json:"platformId"`
	PlatformName      string            `json:"platformName"`
	CreatedAt         time.Time         `json:"createdAt"`
	HoldStatus        bool              `json:"holdStatus"`
	FulfillmentStatus string            `json:"fulfillmentStatus"`
	PaymentStatus     string            `json:"paymentStatus"`
	GrandTotal        float64           `json:"grandTotal"`
	TotalPaid         float64           `json:"totalPaid"`
	RemainPayment     float64           `json:"remainPayment"`
	Creator           OrderCreator      `json:"creator"`
	Consultant        OrderConsultant   `json:"consultant"`
	Customer          OrderCustomer     `json:"customer"`
	ShippingInfo      OrderShippingInfo `json:"shippingInfo"`
	ExternalOrderRef  string            `json:"externalOrderRef"`
}

type OrderCaptureLineItem struct {
	LineItemID        string                          `json:"lineItemId"`
	Sku               string                          `json:"sku"`
	SellerSku         string                          `json:"sellerSku"`
	ServiceID         int64                           `json:"serviceId"`
	DisplayName       string                          `json:"displayName"`
	Uom               string                          `json:"uom"`
	Quantity          float64                         `json:"quantity"`
	CancelledQuantity float64                         `json:"cancelledQuantity"`
	SiteID            int64                           `json:"siteId"`
	Serials           []string                        `json:"serials"`
	Warranty          float64                         `json:"warranty"`
	Price             float64                         `json:"price"`
	OriginalPrice     float64                         `json:"originalPrice"`
	RowTotal          float64                         `json:"rowTotal"`
	IsGift            bool                            `json:"isGift"`
	OnDemandDiscount  float64                         `json:"onDemandDiscount"`
	Promotions        []OrderCaptureLineItemPromotion `json:"promotions"`
	IsAdult           bool                            `json:"isAdult"`
}

type OrderCaptureLineItemPromotion struct {
	ID            int    `json:"id"`
	PromotionID   string `json:"promotionId"`
	PromotionJSON string `json:"promotionJson"`
}

type OrderService struct {
	ID         int64  `json:"id"`
	Sku        string `json:"sku"`
	Name       string `json:"name"`
	Fee        int64  `json:"fee"`
	SellerID   int    `json:"sellerId"`
	SellerName string `json:"sellerName"`
}

type OrderPromotion struct {
	PromotionID int    `json:"promotionId"`
	Benefit     string `json:"benefit"`
}

type OrderPayment struct {
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

type OrderRefund struct {
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

type OrderBillingInfo struct {
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

type Shipment struct {
	ShipmentID                  string         `json:"shipmentId"`
	ShipmentCode                string         `json:"shipmentCode"`
	SellerID                    int64          `json:"sellerId"`
	Status                      string         `json:"status"`
	Items                       []ShipmentItem `json:"items"`
	DocumentIds                 []string       `json:"documentIds"`
	FulfillmentChannelID        int64          `json:"fulfillmentChannelId"`
	FulfillmentChannelName      string         `json:"fulfillmentChannelName"`
	IsAuto                      bool           `json:"isAuto"`
	ReplacedShipmentID          string         `json:"replacedShipmentId"`
	OutboundRequestID           string         `json:"outboundRequestId"`
	OutboundRequestAt           int64          `json:"outboundRequestAt"`
	WarehouseProviderExportID   string         `json:"warehouseProviderExportId"`
	WarehouseExportAt           int64          `json:"warehouseExportAt"`
	WarehouseProviderID         int64          `json:"warehouseProviderId"`
	NumberOfPackages            int64          `json:"numberOfPackages"`
	Height                      int64          `json:"height"`
	Length                      int64          `json:"length"`
	Width                       int64          `json:"width"`
	TotalWeight                 int64          `json:"totalWeight"`
	TransportProviderID         string         `json:"transportProviderId"`
	TransportProviderName       string         `json:"transportProviderName"`
	TransportProviderType       string         `json:"transportProviderType"`
	TransportTrackingID         string         `json:"transportTrackingId"`
	AllowTrialOnDelivery        string         `json:"allowTrialOnDelivery"`
	DeliveryMethodCode          string         `json:"deliveryMethodCode"`
	DeliveryRequestCreatedAt    time.Time      `json:"deliveryRequestCreatedAt"`
	DeliveryRequestCreatorIamID string         `json:"deliveryRequestCreatorIamId"`
	DeliveryRequestID           string         `json:"deliveryRequestId"`
	DeliveryRequestStatus       string         `json:"deliveryRequestStatus"`
	DeliveryTypeGroupID         string         `json:"deliveryTypeGroupId"`
	DeliveryTypeID              string         `json:"deliveryTypeId"`
	Description                 string         `json:"description"`
	ExpectDeliveryDueFrom       time.Time      `json:"expectDeliveryDueFrom"`
	ExpectDeliveryDueTo         time.Time      `json:"expectDeliveryDueTo"`
}

type ShipmentItem struct {
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

type ReturnInfo struct {
	ReturnRequestID           int64            `json:"returnRequestId"`
	PlatformID                string           `json:"platformId"`
	ReasonID                  string           `json:"reasonId"`
	Reason                    string           `json:"reason"`
	SiteID                    int              `json:"siteId"`
	WarehouseProviderExportID string           `json:"warehouseProviderExportId"`
	Status                    string           `json:"status"`
	Items                     []ReturnInfoItem `json:"items"`
	CustomerNotify            bool             `json:"customerNotify"`
	CustomerInstruction       string           `json:"customerInstruction"`
	Comment                   string           `json:"comment"`
	CreatedByEmail            string           `json:"createdByEmail"`
	ApprovedByEmail           string           `json:"approvedByEmail"`
	CreatedAt                 time.Time        `json:"createdAt"`
	UpdatedAt                 time.Time        `json:"updatedAt"`
}

type ReturnInfoItem struct {
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

type OrderCreator struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderConsultant struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CustomID  string `json:"customId"`
}

type OrderCustomer struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderShippingInfo struct {
	Name                          string    `json:"name"`
	Phone                         string    `json:"phone"`
	Email                         string    `json:"email"`
	AddressID                     string    `json:"addressId"`
	Address                       string    `json:"address"`
	WardID                        string    `json:"wardId"`
	WardName                      string    `json:"wardName"`
	DistrictID                    string    `json:"districtId"`
	DistrictName                  string    `json:"districtName"`
	ProvinceID                    string    `json:"provinceId"`
	ProvinceName                  string    `json:"provinceName"`
	FullAddress                   string    `json:"fullAddress"`
	Country                       string    `json:"country"`
	Latitude                      string    `json:"latitude"`
	Longitude                     string    `json:"longitude"`
	Note                          string    `json:"note"`
	SiteID                        int       `json:"siteId"`
	CollectionHubPartnerID        string    `json:"collectionHubPartnerId"`
	EstimatedDeliveryDate         time.Time `json:"estimatedDeliveryDate"`
	ScheduledDeliveryDate         time.Time `json:"scheduledDeliveryDate"`
	ScheduledDeliveryTimeSlotFrom string    `json:"scheduledDeliveryTimeSlotFrom"`
	ScheduledDeliveryTimeSlotTo   string    `json:"scheduledDeliveryTimeSlotTo"`
}
