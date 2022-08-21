package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Invoice struct {
	OrderID     string      `db:"order_id"`
	ShipmentID  string      `db:"shipment_id"`
	SellerID    int64       `db:"seller_id"`
	InvoiceInfo InvoiceInfo `db:"invoice_info"`
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
}

type InvoiceInfo struct {
	ID                                  string             `json:"id"`
	Type                                string             `json:"type"`
	OrderID                             string             `json:"orderId"`
	ShipmentSetID                       string             `json:"shipmentSetId"`
	ShipmentID                          string             `json:"shipmentId"`
	SellerID                            int64              `json:"sellerId"`
	TotalPretaxAmount                   int64              `json:"totalPretaxAmount"`
	TotalVatNAmount                     int64              `json:"totalVatNAmount"`
	TotalVat0Amount                     int64              `json:"totalVat0Amount"`
	TotalVat5Amount                     int64              `json:"totalVat5Amount"`
	TotalVat10Amount                    int64              `json:"totalVat10Amount"`
	GrandTotal                          int64              `json:"grandTotal"`
	BaseTotalAfterTaxAmount             int64              `json:"baseTotalAfterTaxAmount"`
	TotalServiceFeeAfterTaxAmount       int64              `json:"totalServiceFeeAfterTaxAmount"`
	TotalOnDemandDiscountAfterTaxAmount int64              `json:"totalOnDemandDiscountAfterTaxAmount"`
	TotalSellerPromotionAfterTaxAmount  int64              `json:"totalSellerPromotionAfterTaxAmount"`
	TotalPromotionAfterTaxAmount        int64              `json:"totalPromotionAfterTaxAmount"`
	SellerTotalAfterTaxAmount           int64              `json:"sellerTotalAfterTaxAmount"`
	InvoiceItems                        []InvoiceItem      `json:"invoiceItems"`
	Payments                            []InvoicePayment   `json:"payments"`
	PretaxNPromoAmount                  int64              `json:"pretaxNPromoAmount"`
	Pretax0PromoAmount                  int64              `json:"pretax0PromoAmount"`
	Pretax5PromoAmount                  int64              `json:"pretax5PromoAmount"`
	Pretax10PromoAmount                 int64              `json:"pretax10PromoAmount"`
	GrandTotalInWords                   string             `json:"grandTotalInWords"`
	PaidAmount                          int64              `json:"paidAmount"`
	CreditAmount                        int64              `json:"creditAmount"`
	DueAfterDelivery                    int64              `json:"dueAfterDelivery"`
	CodAmount                           int64              `json:"codAmount"`
	TotalAfterTaxOrderDiscount          int64              `json:"totalAfterTaxOrderDiscount"`
	BillingInfo                         InvoiceBillingInfo `json:"billingInfo"`
}

func (i InvoiceInfo) Value() (driver.Value, error) {
	valueString, err := json.Marshal(i)
	return string(valueString), err
}

func (i *InvoiceInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &i)
}

type InvoiceItem struct {
	LineItemID                          string                             `json:"lineItemId"`
	Type                                string                             `json:"type"`
	Sku                                 string                             `json:"sku"`
	Uom                                 string                             `json:"uom"`
	Name                                string                             `json:"name"`
	UnitPretaxPrice                     int64                              `json:"unitPretaxPrice"`
	UnitAfterTaxPrice                   int64                              `json:"unitAfterTaxPrice"`
	Quantity                            int64                              `json:"quantity"`
	TotalAfterTaxAmount                 int64                              `json:"totalAfterTaxAmount"`
	TotalPretaxAmount                   int64                              `json:"totalPretaxAmount"`
	Vat                                 int64                              `json:"vat"`
	VatString                           string                             `json:"vatString"`
	IsGift                              bool                               `json:"isGift"`
	BaseTotalAfterTaxAmount             int64                              `json:"baseTotalAfterTaxAmount"`
	PriceAdjustmentPieces               []InvoiceItemPriceAdjustmentPieces `json:"priceAdjustmentPieces"`
	TotalServiceFeeAfterTaxAmount       int                                `json:"totalServiceFeeAfterTaxAmount"`
	TotalOnDemandDiscountAfterTaxAmount int                                `json:"totalOnDemandDiscountAfterTaxAmount"`
	TotalSellerPromotionAfterTaxAmount  int                                `json:"totalSellerPromotionAfterTaxAmount"`
	TotalPromotionAfterTaxAmount        int                                `json:"totalPromotionAfterTaxAmount"`
	SellerTotalAfterTaxAmount           int                                `json:"sellerTotalAfterTaxAmount"`
}

type InvoiceItemPriceAdjustmentPieces struct {
	ID             string `json:"id"`
	Type           string `json:"type"`
	AfterTaxAmount int64  `json:"afterTaxAmount"`
}

type InvoicePayment struct {
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

type InvoiceBillingInfo struct {
	Name              string    `json:"name"`
	Address           string    `json:"address"`
	Email             string    `json:"email"`
	TaxCode           int64     `json:"taxCode"`
	Phone             int       `json:"phone"`
	Type              int       `json:"type"`
	Note              string    `json:"note"`
	ExpectedIssueDate time.Time `json:"expectedIssueDate"`
}
