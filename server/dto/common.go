package dto

import "time"

type ListOrdersResponseData struct {
	Orders     []ListOrdersBaseOrder `json:"orders"`
	NextOffset int                   `json:"nextOffset"`
}

type ListOrdersBaseOrder struct {
	OrderID           string                   `json:"orderId"`
	ReplacedOrderID   int64                    `json:"replacedOrderId"`
	TerminalID        int                      `json:"terminalId"`
	TerminalName      string                   `json:"terminalName"`
	PlatformID        int64                    `json:"platformId"`
	PlatformName      string                   `json:"platformName"`
	CreatedAt         int                      `json:"createdAt"`
	HoldStatus        bool                     `json:"holdStatus"`
	FulfillmentStatus string                   `json:"fulfillmentStatus"`
	PaymentStatus     string                   `json:"paymentStatus"`
	GrandTotal        int                      `json:"grandTotal"`
	TotalPaid         int                      `json:"totalPaid"`
	RemainPayment     int                      `json:"remainPayment"`
	Creator           OrderDetailsCreator      `json:"creator"`
	Consultant        OrderDetailsConsultant   `json:"consultant"`
	Customer          OrderDetailsCustomer     `json:"customer"`
	ShippingInfo      OrderDetailsShippingInfo `json:"shippingInfo"`
	ExternalOrderRef  string                   `json:"externalOrderRef"`
}

type OrderDetailsCreator struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderDetailsConsultant struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CustomID  string `json:"customId"`
}

type OrderDetailsCustomer struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type OrderDetailsShippingInfo struct {
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
