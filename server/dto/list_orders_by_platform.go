package dto

import (
	"fmt"
	"strings"

	slicehelper "github.com/dokyan1989/postgresdemo/helper/slice"
	"gopkg.in/guregu/null.v4"
)

type ListOrdersByPlatformRequest struct {
	PlatformID         int64     `json:"platformId" schema:"platformId"`
	SellerID           int64     `json:"sellerId" schema:"sellerId"`
	TerminalID         int64     `json:"terminalId" schema:"terminalId"`
	SiteID             int64     `json:"siteId" schema:"siteId"`
	CreatorID          string    `json:"creatorId" schema:"creatorId"`
	ConsultantID       string    `json:"consultantId" schema:"consultantId"`
	HoldStatus         null.Bool `json:"holdStatus" schema:"holdStatus"`
	FulfillmentStatus  string    `json:"fulfillmentStatus" schema:"fulfillmentStatus"`
	PaymentStatus      string    `json:"paymentStatus" schema:"paymentStatus"`
	ConfirmationStatus string    `json:"confirmationStatus" schema:"confirmationStatus"`
	Customer           string    `json:"customer" schema:"customer"`
	OrderID            string    `json:"orderId" schema:"orderId"`
	CreatedAtGte       null.Time `json:"createdAtGte" schema:"createdAtGte"`
	CreatedAtLte       null.Time `json:"createdAtLte" schema:"createdAtLte"`
	SortBy             string    `json:"sortBy" schema:"sortBy"`
	SortOrder          string    `json:"sortOrder" schema:"sortOrder"`
	Limit              uint64    `json:"limit" schema:"limit"`
	Offset             uint64    `json:"offset" schema:"offset"`
}

func (r *ListOrdersByPlatformRequest) Validate() error {
	validSortByValues := []string{"created_at", "updated_at"}
	if !slicehelper.ContainsString(validSortByValues, strings.ToLower(r.SortBy)) {
		return fmt.Errorf("sortBy value must be one of the following: %+v", validSortByValues)
	}

	validSortOrderValues := []string{"desc", "asc"}
	if !slicehelper.ContainsString(validSortOrderValues, strings.ToLower(r.SortOrder)) {
		return fmt.Errorf("sortOrder value must be one of the following: %+v", validSortOrderValues)
	}

	if !r.CreatedAtGte.Valid || !r.CreatedAtLte.Valid {
		return fmt.Errorf("createdAtGte and createdAtLte cannot be empty")
	}

	return nil
}

type ListOrdersByPlatformResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    ListOrdersResponseData `json:"data"`
}
