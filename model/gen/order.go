package gen

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dokyan1989/postgresdemo/helper/random"
	"github.com/dokyan1989/postgresdemo/model"
	"github.com/google/uuid"
)

type GenOrderOption func(o *model.Order)

func GenOrderWithID(id string) GenOrderOption {
	return func(o *model.Order) {
		if id != "" {
			o.ID = id
		}
	}
}

func GenOrder(opts ...GenOrderOption) *model.Order {
	now := time.Now()
	customerPhone := gofakeit.Phone()
	customerName := gofakeit.Name()
	customerEmail := gofakeit.Email()
	shippingInfoPhone := gofakeit.Phone()

	order := &model.Order{
		ID:                 uuid.NewString(),
		PaymentStatus:      GenPaymentStatus(),
		ConfirmationStatus: GenConfirmationStatus(),
		CustomerPhone:      customerPhone,
		CustomerName:       customerName,
		CustomerEmail:      customerEmail,
		ShippingInfoPhone:  shippingInfoPhone,
		TerminalId:         int64(random.Int(1, 10)),
		PlatformId:         int64(random.Int(1, 10)),
		CreatorId:          uuid.NewString(),
		ConsultantId:       uuid.NewString(),
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	for _, opt := range opts {
		opt(order)
	}

	return order
}

func GenOrderRawData(order *model.Order) *model.OrderRawData {
	if order == nil {
		return nil
	}

	orderRawData := &model.OrderRawData{
		ID:        order.ID,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

	// gen order info
	var orderInfo model.OrderInfo
	gofakeit.Struct(&orderInfo)
	orderInfo.ID = order.ID
	orderInfo.ConfirmationStatus = order.ConfirmationStatus
	orderInfo.PaymentStatus = order.PaymentStatus

	// gen return info
	rrNum := random.Int(0, 3)
	returnInfo := make(model.ReturnInfo, rrNum)
	for i := 0; i < rrNum; i++ {
		var returnRequest model.ReturnRequest
		gofakeit.Struct(&returnRequest)

		returnRequest.OrderID = order.ID

		returnInfo[i] = returnRequest
	}

	// set order info and return info to order raw data
	orderRawData.OrderInfo = orderInfo
	orderRawData.ReturnInfo = returnInfo

	return orderRawData
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
