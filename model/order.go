package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Order struct {
	ID                 string        `db:"id"`
	FulfillmentStatus  string        `db:"fulfillment_status"`
	PaymentStatus      string        `db:"payment_status"`
	ConfirmationStatus string        `db:"confirmation_status"`
	HoldStatus         bool          `db:"hold_status"`
	CustomerPhone      string        `db:"customer_phone"`
	CustomerName       string        `db:"customer_name"`
	CustomerEmail      string        `db:"customer_email"`
	ShippingInfoPhone  string        `db:"shipping_info_phone"`
	TerminalId         int64         `db:"terminal_id"`
	PlatformId         int64         `db:"platform_id"`
	CreatorId          string        `db:"creator_id"`
	ConsultantId       string        `db:"consultant_id"`
	SiteIds            pq.Int64Array `db:"site_ids"`
	CreatedAt          time.Time     `db:"created_at"`
	UpdatedAt          time.Time     `db:"updated_at"`
}

type OrderRawData struct {
	ID         string     `db:"id"`
	OrderInfo  OrderInfo  `db:"order_info"`
	ReturnInfo ReturnInfo `db:"return_info"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
}

/**
|-------------------------------------------------------------------------
| Order Info
|-----------------------------------------------------------------------*/
type OrderInfo struct {
	ID                      string                       `json:"id"`
	OrderToken              string                       `json:"orderToken"`
	GrandTotal              float64                      `json:"grandTotal"`
	RemainPayment           float64                      `json:"remainPayment"`
	QrPromotionAmount       float64                      `json:"qrPromotionAmount"`
	TotalFee                float64                      `json:"totalFee"`
	PlatformID              int64                        `json:"platformId"`
	TerminalCode            string                       `json:"terminalCode"`
	TerminalID              int64                        `json:"terminalId"`
	CustomerInfo            OrderCustomerInfo            `json:"customerInfo"`
	Creator                 OrderCreator                 `json:"creator"`
	Items                   []OrderItem                  `json:"items"`
	BillingInfo             OrderBillingInfo             `json:"billingInfo"`
	DeliveryInfo            OrderDeliveryInfo            `json:"deliveryInfo"`
	Consultant              OrderConsultant              `json:"consultant"`
	CreatedAt               time.Time                    `json:"createdAt"`
	UpdatedAt               time.Time                    `json:"updatedAt"`
	Note                    string                       `json:"note"`
	PaymentStatus           string                       `json:"paymentStatus"`
	HoldStatus              bool                         `json:"holdStatus"`
	ConfirmationStatus      string                       `json:"confirmationStatus"`
	Promotions              []Promotion                  `json:"promotions"`
	OrderPromotions         []OrderPromotion             `json:"orderPromotions"`
	PromotionDiscount       float64                      `json:"promotionDiscount"`
	OnDemandDiscount        float64                      `json:"onDemandDiscount"`
	TotalDiscount           float64                      `json:"totalDiscount"`
	UpdatedBy               string                       `json:"updatedBy"`
	ReplaceOrderID          string                       `json:"replaceOrderId"`
	ExternalOrderRef        string                       `json:"externalOrderRef"`
	SendEmailToBuyer        bool                         `json:"sendEmailToBuyer"`
	IsHandover              bool                         `json:"isHandover"`
	TotalPaid               float64                      `json:"totalPaid"`
	Services                []OrderService               `json:"services"`
	Payments                []OrderPayment               `json:"payments"`
	DiscountApproval        OrderDiscountApproval        `json:"discountApproval"`
	DeferredPaymentApproval OrderDeferredPaymentApproval `json:"deferredPaymentApproval"`
}

func (o OrderInfo) Value() (driver.Value, error) {
	valueString, err := json.Marshal(o)
	return string(valueString), err
}

func (o *OrderInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}

type OrderCustomerInfo struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderCreator struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderItem struct {
	LineItemID        string               `json:"lineItemId"`
	Sku               string               `json:"sku"`
	SellerSku         string               `json:"sellerSku"`
	ServiceID         int64                `json:"serviceId"`
	DisplayName       string               `json:"displayName"`
	Uom               string               `json:"uom"`
	Quantity          float64              `json:"quantity"`
	CancelledQuantity float64              `json:"cancelledQuantity"`
	SiteID            int64                `json:"siteId"`
	Serials           []string             `json:"serials"`
	Warranty          float64              `json:"warranty"`
	Price             float64              `json:"price"`
	OriginalPrice     float64              `json:"originalPrice"`
	RowTotal          float64              `json:"rowTotal"`
	IsGift            bool                 `json:"isGift"`
	OnDemandDiscount  float64              `json:"onDemandDiscount"`
	Promotions        []OrderItemPromotion `json:"promotions"`
	IsAdult           bool                 `json:"isAdult"`
}

type OrderItemPromotion struct {
	ID            int    `json:"id"`
	PromotionID   string `json:"promotionId"`
	PromotionJSON string `json:"promotionJson"`
}

type OrderBillingInfo struct {
	Name              string    `json:"name"`
	Address           string    `json:"address"`
	Email             string    `json:"email"`
	TaxCode           int64     `json:"taxCode"`
	Phone             int       `json:"phone"`
	Type              int       `json:"type"`
	Note              string    `json:"note"`
	ExpectedIssueDate time.Time `json:"expectedIssueDate"`
}

type OrderDeliveryInfo struct {
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

type OrderConsultant struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CustomID  string `json:"customId"`
}

type OrderPromotion struct {
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
	GroupID    int    `json:"groupId"`
}

type OrderPayment struct {
	MethodCode             string `json:"methodCode"`
	Amount                 int64  `json:"amount"`
	TransactionCode        string `json:"transactionCode"`
	TotalPaid              int64  `json:"totalPaid"`
	MerchantCode           string `json:"merchantCode"`
	PartnerTransactionCode string `json:"partnerTransactionCode"`
	PaymentService         string `json:"paymentService"`
	CashierID              string `json:"cashierId"`
}

type OrderDiscountApproval struct {
	ApproverID string `json:"approverId"`
	Amount     int64  `json:"amount"`
	ReasonID   int64  `json:"reasonId"`
	Reason     string `json:"reason"`
	Status     string `json:"status"`
	Note       string `json:"note"`
}

type OrderDeferredPaymentApproval struct {
	Amount                         int64  `json:"amount"`
	DeferredPaymentPeriod          int64  `json:"deferredPaymentPeriod"`
	DeferredPaymentRequireApproval bool   `json:"deferredPaymentRequireApproval"`
	ApproverID                     string `json:"approverId"`
	Status                         string `json:"status"`
	Note                           string `json:"note"`
}

type Promotion struct {
	ID               int                     `json:"id"`
	PromotionID      string                  `json:"promotionId"`
	Type             string                  `json:"type"`
	ApplyType        string                  `json:"applyType"`
	IsDefault        bool                    `json:"isDefault"`
	Coupon           string                  `json:"coupon"`
	ApplyOn          []PromotionApplyOn      `json:"applyOn"`
	Discount         int                     `json:"discount"`
	OriginalDiscount int                     `json:"originalDiscount"`
	DiscountItems    []PromotionDiscountItem `json:"discountItems"`
	Gifts            []PromotionGift         `json:"gifts"`
	RemovedGifts     []PromotionRemovedGift  `json:"removedGifts"`
	Quantity         int                     `json:"quantity"`
	Voucher          PromotionVoucher        `json:"voucher"`
}

type PromotionApplyOn struct {
	LineItemID       string `json:"lineItemId"`
	ServiceID        int64  `json:"serviceId"`
	Quantity         int    `json:"quantity"`
	ComboConditionID string `json:"comboConditionId"`
	Sku              string `json:"sku"`
}

type PromotionDiscountItem struct {
	LineItemID string `json:"lineItemId"`
	Sku        string `json:"sku"`
	Quantity   int64  `json:"quantity"`
}

type PromotionGift struct {
	LineItemID string `json:"lineItemId"`
	Sku        string `json:"sku"`
	Quantity   int    `json:"quantity"`
	Name       string `json:"name"`
}

type PromotionRemovedGift struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type PromotionVoucher struct {
	Quantity int `json:"quantity"`
}

/**
|-------------------------------------------------------------------------
| Return Info
|-----------------------------------------------------------------------*/
type ReturnInfo []ReturnRequest

func (r ReturnInfo) Value() (driver.Value, error) {
	valueString, err := json.Marshal(r)
	return string(valueString), err
}

func (r *ReturnInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &r)
}

type ReturnRequest struct {
	OrderID                   string              `json:"orderId"`
	ReturnRequestID           int64               `json:"returnRequestId"`
	PlatformID                string              `json:"platformId"`
	ReasonID                  string              `json:"reasonId"`
	Reason                    string              `json:"reason"`
	SiteID                    int                 `json:"siteId"`
	WarehouseProviderExportID string              `json:"warehouseProviderExportId"`
	Status                    string              `json:"status"`
	Items                     []ReturnRequestItem `json:"items"`
	CustomerNotify            bool                `json:"customerNotify"`
	CustomerInstruction       string              `json:"customerInstruction"`
	Comment                   string              `json:"comment"`
	CreatedByEmail            string              `json:"createdByEmail"`
	ApprovedByEmail           string              `json:"approvedByEmail"`
	CreatedAt                 time.Time           `json:"createdAt"`
	UpdatedAt                 time.Time           `json:"updatedAt"`
}

type ReturnRequestItem struct {
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
