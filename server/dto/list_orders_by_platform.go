package dto

type ListOrdersByPlatformRequest struct {
	PlatformID         int    `json:"platformId" schema:"platformId"`
	SellerID           int    `json:"sellerId" schema:"sellerId"`
	FulfillmentStatus  string `json:"fulfillmentStatus" schema:"fulfillmentStatus"`
	TerminalID         int    `json:"terminalId" schema:"terminalId"`
	CreatorID          string `json:"creatorId" schema:"creatorId"`
	ConsultantID       string `json:"consultantId" schema:"consultantId"`
	Customer           string `json:"customer" schema:"customer"`
	HoldStatus         string `json:"holdStatus" schema:"holdStatus"`
	PaymentStatus      string `json:"paymentStatus" schema:"paymentStatus"`
	ConfirmationStatus string `json:"confirmationStatus" schema:"confirmationStatus"`
	OrderID            string `json:"orderId" schema:"orderId"`
	Limit              int    `json:"limit" schema:"limit"`
	Offset             int    `json:"offset" schema:"offset"`
	SortBy             string `json:"sortBy" schema:"sortBy"`
	SortOrder          string `json:"sortOrder" schema:"sortOrder"`
	CreatedAtGte       string `json:"createdAtGte" schema:"createdAtGte"`
	CreatedAtLlte      string `json:"createdAtLlte" schema:"createdAtLlte"`
}

type ListOrdersByPlatformResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    ListOrdersResponseData `json:"data"`
}
