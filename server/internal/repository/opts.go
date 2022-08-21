package repository

import "gopkg.in/guregu/null.v4"

type ListOrdersOptions struct {
	PlatformID         int64
	SellerID           int64
	TerminalID         int64
	SiteID             int64
	CreatorID          string
	ConsultantID       string
	HoldStatus         null.Bool
	FulfillmentStatus  string
	PaymentStatus      string
	ConfirmationStatus string
	Customer           string
	OrderID            string
	CreatedAtGte       null.Time
	CreatedAtLte       null.Time
	SortBy             string
	SortOrder          string
	Limit              uint64
	Offset             uint64
}
