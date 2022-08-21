package gen

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dokyan1989/postgresdemo/helper/random"
	"github.com/dokyan1989/postgresdemo/model"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func GenOrder(orderRawData *model.OrderRawData) *model.Order {
	if orderRawData == nil {
		return nil
	}

	siteIds := make([]int64, 0)
	exists := make(map[int64]bool, 0)
	for _, item := range orderRawData.OrderInfo.Items {
		if _, ok := exists[item.SiteID]; ok {
			continue
		}
		siteIds = append(siteIds, item.SiteID)
		exists[item.SiteID] = true
	}
	order := &model.Order{
		ID:                 orderRawData.ID,
		FulfillmentStatus:  GenFulfillmentStatus(),
		PaymentStatus:      orderRawData.OrderInfo.PaymentStatus,
		HoldStatus:         orderRawData.OrderInfo.HoldStatus,
		ConfirmationStatus: orderRawData.OrderInfo.ConfirmationStatus,
		CustomerPhone:      orderRawData.OrderInfo.CustomerInfo.Phone,
		CustomerName:       orderRawData.OrderInfo.CustomerInfo.Name,
		CustomerEmail:      orderRawData.OrderInfo.CustomerInfo.Email,
		ShippingInfoPhone:  orderRawData.OrderInfo.DeliveryInfo.Phone,
		TerminalId:         orderRawData.OrderInfo.TerminalID,
		PlatformId:         orderRawData.OrderInfo.PlatformID,
		CreatorId:          orderRawData.OrderInfo.Creator.ID,
		ConsultantId:       orderRawData.OrderInfo.Consultant.ID,
		SiteIds:            siteIds,
		CreatedAt:          orderRawData.CreatedAt,
		UpdatedAt:          orderRawData.UpdatedAt,
	}

	return order
}

func GenOrderRawData() (*model.OrderRawData, error) {
	id, err := sf.NextID()
	if err != nil {
		return nil, err
	}

	orderRawData := &model.OrderRawData{
		ID:        fmt.Sprintf("%d", id),
		CreatedAt: gofakeit.DateRange(time.Now().Add((-3)*OneYear), time.Now()),
		UpdatedAt: gofakeit.DateRange(time.Now().Add((-3)*OneYear), time.Now()),
	}

	// gen order info
	var orderInfo model.OrderInfo
	gofakeit.Struct(&orderInfo)

	orderInfo.ID = orderRawData.ID
	orderInfo.ConfirmationStatus = GenConfirmationStatus()
	orderInfo.PaymentStatus = GenPaymentStatus()
	orderInfo.HoldStatus = gofakeit.Bool()
	orderInfo.PlatformID = int64(random.Int(1, 10))
	orderInfo.TerminalID = int64(random.Int(1, 10))
	orderInfo.CustomerInfo.Phone = gofakeit.Phone()
	orderInfo.CustomerInfo.Email = gofakeit.Email()
	orderInfo.CustomerInfo.Name = gofakeit.Name()
	orderInfo.DeliveryInfo.Phone = gofakeit.Phone()
	orderInfo.Creator.ID = GenCreatorOrConsultantID()
	orderInfo.Consultant.ID = GenCreatorOrConsultantID()
	orderInfo.CreatedAt = orderRawData.CreatedAt
	orderInfo.UpdatedAt = orderRawData.UpdatedAt

	for i := range orderInfo.Items {
		orderInfo.Items[i].SiteID = int64(random.Int(1, 10))
	}

	// gen return info
	rrNum := random.Int(0, 3)
	returnInfo := make(model.ReturnInfo, rrNum)
	for i := 0; i < rrNum; i++ {
		var returnRequest model.ReturnRequest
		gofakeit.Struct(&returnRequest)

		returnRequest.OrderID = orderRawData.ID
		returnInfo[i] = returnRequest
	}

	// set order info and return info to order raw data
	orderRawData.OrderInfo = orderInfo
	orderRawData.ReturnInfo = returnInfo

	return orderRawData, nil
}

func GenPaymentStatus() string {
	values := []string{
		"PENDING",
		"PARTIALLY_PAID",
		"FULLY_PAID",
	}

	i := random.Int(0, len(values)-1)
	return values[i]
}

func GenConfirmationStatus() string {
	values := []string{
		"PENDING",
		"ACTIVE",
		"CANCELLED",
	}

	i := random.Int(0, len(values)-1)
	return values[i]
}

func GenFulfillmentStatus() string {
	values := []string{
		"NEW",
		"WAITING_FOR_MANUAL_CONFIRMATION",
		"PROVIDER_TRANSFERRING",
		"PROVIDER_TRANSFERRED",
		"WAITING_FOR_PRODUCT",
		"PICKED",
		"PACKED",
		"DELIVERING",
		"DELIVERED",
		"PARTIAL_DELIVERED",
		"CANCELLED",
		"RETURNED",
		"PARTIAL_CANCELLED",
	}

	i := random.Int(0, len(values)-1)
	return values[i]
}

func GenCreatorOrConsultantID() string {
	values := []string{
		"EKRLIU9554",
		"ILH6J5GB58",
		"VSS0GYI4Z5",
		"U0OSVBZNTC",
		"D6ITVD9KLL",
		"ZNCZ7H5KGO",
		"MY0ZP6S2IP",
		"TNJ37RK8TJ",
		"PLKCLFB477",
		"OCW5WCPBPU",
		"53WL2FWPF4",
		"H7GDCB514H",
		"K75W121F4N",
		"HL8ZD41YJB",
		"GMIKGH1Z1A",
		"OVEKLJC5FO",
		"UU0J4HIN65",
		"NLY89UG67D",
		"MZWHROF4MH",
		"EKFZXRGC0L",
	}

	i := random.Int(0, len(values)-1)
	return values[i]
}
