package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Order struct {
	ID                 string    `db:"id"`
	FulfillmentStatus  string    `db:"fulfillment_status"`
	PaymentStatus      string    `db:"payment_status"`
	ConfirmationStatus string    `db:"confirmation_status"`
	CustomerPhone      string    `db:"customer_phone"`
	CustomerName       string    `db:"customer_name"`
	CustomerEmail      string    `db:"customer_email"`
	ShippingInfoPhone  string    `db:"shipping_info_phone"`
	TerminalId         int64     `db:"terminal_id"`
	PlatformId         int64     `db:"platform_id"`
	CreatorId          string    `db:"creator_id"`
	ConsultantId       string    `db:"consultant_id"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
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
	ID                      string                    `json:"id"`
	OrderToken              string                    `json:"orderToken"`
	GrandTotal              int64                     `json:"grandTotal"`
	RemainPayment           int                       `json:"remainPayment"`
	QrPromotionAmount       int                       `json:"qrPromotionAmount"`
	TotalFee                int64                     `json:"totalFee"`
	PlatformID              int64                     `json:"platformId"`
	TerminalCode            string                    `json:"terminalCode"`
	TerminalID              int                       `json:"terminalId"`
	CustomerInfo            OICustomerInfo            `json:"customerInfo"`
	Creator                 OICreator                 `json:"creator"`
	Items                   []OIItem                  `json:"items"`
	BillingInfo             OIBillingInfo             `json:"billingInfo"`
	DeliveryInfo            OIDeliveryInfo            `json:"deliveryInfo"`
	Consultant              OIConsultant              `json:"consultant"`
	CreatedAt               string                    `json:"createdAt"`
	UpdatedAt               string                    `json:"updatedAt"`
	Note                    string                    `json:"note"`
	PaymentStatus           string                    `json:"paymentStatus"`
	HoldStatus              bool                      `json:"holdStatus"`
	ConfirmationStatus      string                    `json:"confirmationStatus"`
	OrderPromotion          []OIOrderPromotion        `json:"orderPromotion"`
	PromotionDiscount       int                       `json:"promotionDiscount"`
	OnDemandDiscount        int64                     `json:"onDemandDiscount"`
	TotalDiscount           int64                     `json:"totalDiscount"`
	UpdatedBy               string                    `json:"updatedBy"`
	ReplaceOrderID          string                    `json:"replaceOrderId"`
	ExternalOrderRef        string                    `json:"externalOrderRef"`
	SendEmailToBuyer        bool                      `json:"sendEmailToBuyer"`
	IsHandover              bool                      `json:"isHandover"`
	TotalPaid               int                       `json:"totalPaid"`
	Services                []OIService               `json:"services"`
	Payments                []OIPayment               `json:"payments"`
	DiscountApproval        OIDiscountApproval        `json:"discountApproval"`
	DeferredPaymentApproval OIDeferredPaymentApproval `json:"deferredPaymentApproval"`
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

type OICustomerInfo struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	CustomID  string `json:"customId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OICreator struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	CustomID  string `json:"customId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OIItem struct {
	LineItemID        string            `json:"lineItemId"`
	Sku               string            `json:"sku"`
	SellerSku         string            `json:"sellerSku"`
	ServiceID         int64             `json:"serviceId"`
	DisplayName       string            `json:"displayName"`
	Uom               string            `json:"uom"`
	Quantity          int               `json:"quantity"`
	CancelledQuantity int               `json:"cancelledQuantity"`
	SiteID            int               `json:"siteId"`
	Serials           []string          `json:"serials"`
	Warranty          int               `json:"warranty"`
	Price             int               `json:"price"`
	OriginalPrice     int               `json:"originalPrice"`
	RowTotal          int               `json:"rowTotal"`
	IsGift            bool              `json:"isGift"`
	OnDemandDiscount  int               `json:"onDemandDiscount"`
	Promotions        []OIItemPromotion `json:"promotions"`
	IsAdult           bool              `json:"isAdult"`
}

type OIBillingInfo struct {
	Name              string    `json:"name"`
	Address           string    `json:"address"`
	Email             string    `json:"email"`
	TaxCode           string    `json:"taxCode"`
	Phone             string    `json:"phone"`
	Type              string    `json:"type"`
	Note              string    `json:"note"`
	ExpectedIssueDate time.Time `json:"expectedIssueDate"`
}

type OIDeliveryInfo struct {
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
	Lattitude                     string    `json:"lattitude"`
	Longitude                     string    `json:"longitude"`
	Note                          string    `json:"note"`
	SiteID                        int       `json:"siteId"`
	CollectionHubPartnerID        string    `json:"collectionHubPartnerId"`
	EstimatedDeliveryDate         time.Time `json:"estimatedDeliveryDate"`
	ScheduledDeliveryDate         time.Time `json:"scheduledDeliveryDate"`
	ScheduledDeliveryTimeSlotFrom string    `json:"scheduledDeliveryTimeSlotFrom"`
	ScheduledDeliveryTimeSlotTo   string    `json:"scheduledDeliveryTimeSlotTo"`
}

type OIConsultant struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	CustomID  string `json:"customId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OIOrderPromotion struct {
	PromotionID int    `json:"promotionId"`
	Benefit     string `json:"benefit"`
}

type OIService struct {
	ID         int64  `json:"id"`
	Sku        string `json:"sku"`
	Name       string `json:"name"`
	Fee        int64  `json:"fee"`
	SellerID   int    `json:"sellerId"`
	SellerName string `json:"sellerName"`
	GroupID    int    `json:"groupId"`
}

type OIPayment struct {
	MethodCode             string `json:"methodCode"`
	Amount                 int64  `json:"amount"`
	TransactionCode        string `json:"transactionCode"`
	TotalPaid              int64  `json:"totalPaid"`
	MerchantCode           string `json:"merchantCode"`
	PartnerTransactionCode string `json:"partnerTransactionCode"`
	PaymentService         string `json:"paymentService"`
	CashierID              string `json:"cashierId"`
}

type OIDiscountApproval struct {
	ApproverID string `json:"approverId"`
	Amount     int64  `json:"amount"`
	ReasonID   int64  `json:"reasonId"`
	Reason     string `json:"reason"`
	Status     string `json:"status"`
	Note       string `json:"note"`
}

type OIDeferredPaymentApproval struct {
	Amount                         int64  `json:"amount"`
	DeferredPaymentPeriod          int64  `json:"deferredPaymentPeriod"`
	DeferredPaymentRequireApproval bool   `json:"deferredPaymentRequireApproval"`
	ApproverID                     string `json:"approverId"`
	Status                         string `json:"status"`
	Note                           string `json:"note"`
}

type OIItemPromotion struct {
	ID               int                           `json:"id"`
	PromotionID      string                        `json:"promotionId"`
	Type             string                        `json:"type"`
	ApplyType        string                        `json:"applyType"`
	ApplyOn          []OIItemPromotionApplyOn      `json:"applyOn"`
	Discount         int                           `json:"discount"`
	OriginalDiscount int                           `json:"originalDiscount"`
	DiscountItems    []OIItemPromotionDiscountItem `json:"discountItems"`
	Gifts            []OIItemPromotionGift         `json:"gifts"`
	RemovedGifts     []OIItemPromotionRemovedGift  `json:"removedGifts"`
	Quantity         int                           `json:"quantity"`
	Voucher          OIItemPromotionVoucher        `json:"voucher"`
}

type OIItemPromotionApplyOn struct {
	LineItemID       string `json:"lineItemId"`
	ServiceID        int64  `json:"serviceId"`
	Quantity         int    `json:"quantity"`
	ComboConditionID string `json:"comboConditionId"`
	Sku              string `json:"sku"`
}

type OIItemPromotionDiscountItem struct {
	LineItemID string `json:"lineItemId"`
	Sku        string `json:"sku"`
	Quantity   int64  `json:"quantity"`
}

type OIItemPromotionGift struct {
	LineItemID string `json:"lineItemId"`
	Sku        string `json:"sku"`
	Quantity   int    `json:"quantity"`
	Name       string `json:"name"`
}

type OIItemPromotionRemovedGift struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

type OIItemPromotionVoucher struct {
	Quantity     int    `json:"quantity"`
	Coupon       string `json:"coupon"`
	SellerIds    []int  `json:"sellerIds"`
	ComboApplyOn string `json:"comboApplyOn"`
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
	OrderID                   string    `json:"orderId"`
	ReturnRequestID           int64     `json:"returnRequestId"`
	PlatformID                string    `json:"platformId"`
	ReasonID                  string    `json:"reasonId"`
	Reason                    string    `json:"reason"`
	SiteID                    int       `json:"siteId"`
	WarehouseProviderExportID string    `json:"warehouseProviderExportId"`
	Status                    string    `json:"status"`
	Items                     []RRItem  `json:"items"`
	CustomerNotify            bool      `json:"customerNotify"`
	CustomerInstruction       string    `json:"customerInstruction"`
	Comment                   string    `json:"comment"`
	CreatedByEmail            string    `json:"createdByEmail"`
	ApprovedByEmail           string    `json:"approvedByEmail"`
	CreatedAt                 time.Time `json:"createdAt"`
	UpdatedAt                 time.Time `json:"updatedAt"`
}

type RRItem struct {
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
