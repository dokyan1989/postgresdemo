package transformer

import (
	"github.com/dokyan1989/postgresdemo/model"
	"github.com/dokyan1989/postgresdemo/server/dto"
)

func (dt *DtoTransformer) ToBaseSaleOrders(orders []model.Order, ordersRawData []model.OrderRawData) []dto.BaseSaleOrder {
	ordersMap := make(map[string]model.Order, 0)
	for _, order := range orders {
		ordersMap[order.ID] = order
	}

	data := make([]dto.BaseSaleOrder, len(ordersRawData))
	for i, ord := range ordersRawData {
		data[i] = dt.ToBaseSaleOrder(ordersMap[ord.ID], ord)
	}

	return data
}

func (dt *DtoTransformer) ToBaseSaleOrder(order model.Order, orderRawData model.OrderRawData) dto.BaseSaleOrder {
	return dto.BaseSaleOrder{
		OrderID:           orderRawData.ID,
		ReplacedOrderID:   orderRawData.OrderInfo.ReplaceOrderID,
		TerminalID:        orderRawData.OrderInfo.TerminalID,
		TerminalName:      "",
		PlatformID:        orderRawData.OrderInfo.PlatformID,
		PlatformName:      "",
		CreatedAt:         orderRawData.CreatedAt,
		HoldStatus:        orderRawData.OrderInfo.HoldStatus,
		FulfillmentStatus: order.FulfillmentStatus,
		PaymentStatus:     orderRawData.OrderInfo.PaymentStatus,
		GrandTotal:        orderRawData.OrderInfo.GrandTotal,
		TotalPaid:         orderRawData.OrderInfo.TotalPaid,
		RemainPayment:     orderRawData.OrderInfo.RemainPayment,
		Creator:           dt.ToCreator(orderRawData.OrderInfo.Creator),
		Consultant:        dt.ToConsultant(orderRawData.OrderInfo.Consultant),
		Customer:          dt.ToCustomer(orderRawData.OrderInfo.CustomerInfo),
		ShippingInfo:      dt.ToShippingInfo(orderRawData.OrderInfo.DeliveryInfo),
		ExternalOrderRef:  orderRawData.OrderInfo.ExternalOrderRef,
	}
}
