package transformer

import (
	"strconv"

	"github.com/dokyan1989/postgresdemo/model"
	"github.com/dokyan1989/postgresdemo/server/dto"
	"github.com/spf13/cast"
)

type DtoTransformer struct{}

func (dt *DtoTransformer) ToFullSaleOrder(
	order *model.Order,
	orderRawData *model.OrderRawData,
	shipments []model.Shipment) *dto.FullSaleOrder {

	orderInfo := orderRawData.OrderInfo

	data := &dto.FullSaleOrder{
		TerminalCode:          orderInfo.TerminalCode,
		ReferralCode:          "",
		Status:                "",
		OrderCaptureLineItems: dt.ToOrderCaptureLineItems(orderInfo.Items),
		Services:              dt.ToServices(orderInfo.Services),
		OrderPromotion:        dt.ToOrderPromotion(orderInfo.OrderPromotions),
		TotalDiscount:         orderInfo.TotalDiscount,
		QrPromotionAmount:     orderInfo.QrPromotionAmount,
		Payments:              dt.ToPayments(orderInfo.Payments),
		// Refunds:               []dto.OrderRefund{},
		BillingInfo:           dt.ToBillingInfo(orderInfo.BillingInfo),
		Shipments:             dt.ToShipments(shipments),
		Returns:               dt.ToReturns(orderRawData.ReturnInfo),
		Note:                  orderInfo.Note,
		Source:                "",
		LoyaltyGrantPoint:     0,
		LoyaltyGrantTierPoint: 0,
		OrderID:               orderInfo.ID,
		ReplacedOrderID:       orderInfo.ReplaceOrderID,
		TerminalID:            orderInfo.TerminalID,
		TerminalName:          "",
		PlatformID:            orderInfo.PlatformID,
		PlatformName:          "",
		CreatedAt:             orderInfo.CreatedAt,
		HoldStatus:            orderInfo.HoldStatus,
		FulfillmentStatus:     order.FulfillmentStatus,
		PaymentStatus:         orderInfo.PaymentStatus,
		GrandTotal:            cast.ToFloat64(orderInfo.GrandTotal),
		TotalPaid:             cast.ToFloat64(orderInfo.TotalPaid),
		RemainPayment:         cast.ToFloat64(orderInfo.RemainPayment),
		Creator:               dt.ToCreator(orderInfo.Creator),
		Consultant:            dt.ToConsultant(orderInfo.Consultant),
		Customer:              dt.ToCustomer(orderInfo.CustomerInfo),
		ShippingInfo:          dt.ToShippingInfo(orderInfo.DeliveryInfo),
		ExternalOrderRef:      orderInfo.ExternalOrderRef,
	}

	return data
}

func (dt *DtoTransformer) ToOrderCaptureLineItems(items []model.OrderItem) []dto.OrderCaptureLineItem {
	data := make([]dto.OrderCaptureLineItem, len(items))

	for i, item := range items {
		data[i] = dto.OrderCaptureLineItem{
			LineItemID:        item.LineItemID,
			Sku:               item.Sku,
			SellerSku:         item.SellerSku,
			ServiceID:         item.ServiceID,
			DisplayName:       item.DisplayName,
			Uom:               item.Uom,
			Quantity:          item.Quantity,
			CancelledQuantity: item.CancelledQuantity,
			SiteID:            item.SiteID,
			Serials:           item.Serials,
			Warranty:          item.Warranty,
			Price:             item.Price,
			OriginalPrice:     item.OriginalPrice,
			RowTotal:          item.RowTotal,
			IsGift:            item.IsGift,
			OnDemandDiscount:  item.OnDemandDiscount,
			Promotions:        dt.ToOrderCaptureLineItemPromotions(item.Promotions),
			IsAdult:           item.IsAdult,
		}
	}

	return data
}

func (dt *DtoTransformer) ToOrderCaptureLineItemPromotions(itemPromotions []model.OrderItemPromotion) []dto.OrderCaptureLineItemPromotion {
	data := make([]dto.OrderCaptureLineItemPromotion, len(itemPromotions))

	for i, itemPromotion := range itemPromotions {
		data[i] = dto.OrderCaptureLineItemPromotion{
			ID:            itemPromotion.ID,
			PromotionID:   itemPromotion.PromotionID,
			PromotionJSON: itemPromotion.PromotionJSON,
		}
	}

	return data
}

func (dt *DtoTransformer) ToServices(services []model.OrderService) []dto.OrderService {
	data := make([]dto.OrderService, len(services))

	for i, service := range services {
		data[i] = dto.OrderService{
			ID:         service.ID,
			Sku:        service.Sku,
			Name:       service.Name,
			Fee:        service.Fee,
			SellerID:   service.SellerID,
			SellerName: service.SellerName,
		}
	}

	return data
}

func (dt *DtoTransformer) ToOrderPromotion(promotions []model.OrderPromotion) []dto.OrderPromotion {
	data := make([]dto.OrderPromotion, len(promotions))

	for i, p := range promotions {
		data[i] = dto.OrderPromotion{
			PromotionID: p.ID,
			// Benefit:     "",
		}
	}

	return data
}

func (dt *DtoTransformer) ToPayments(payments []model.OrderPayment) []dto.OrderPayment {
	data := make([]dto.OrderPayment, len(payments))

	for i, payment := range payments {
		data[i] = dto.OrderPayment{
			// PartnerCode:            payment.PartnerCode,
			MethodCode: payment.MethodCode,
			// OrderCode:              payment.OrderCode,
			TransactionCode: payment.TransactionCode,
			// ClientTransactionCode:  payment.ClientTransactionCode,
			Amount: strconv.Itoa(int(payment.Amount)),
			// Status:                 payment.Status,
			// Message:                payment.Message,
			// MerchantMethodCode:     payment.MerchantMethodCode,
			PartnerTransactionCode: payment.PartnerTransactionCode,
			// BankCode:               payment.BankCode,
			// BankAccountNo:          payment.BankAccountNo,
			// PaymentTerm:            payment.PaymentTerm,
			CashierID: payment.CashierID,
			// PaidAt:                 payment.PaidAt,
		}
	}
	return data
}

func (dt *DtoTransformer) ToBillingInfo(billingInfo model.OrderBillingInfo) dto.OrderBillingInfo {
	return dto.OrderBillingInfo{
		// SplitInvoice:              billingInfo.SplitInvoice,
		// SplitInvoiceGrandTotalMax: billingInfo.SplitInvoiceGrandTotalMax,
		// NewInvoiceDate:            billingInfo.NewInvoiceDate,
		Name:              billingInfo.Name,
		Address:           billingInfo.Address,
		Email:             billingInfo.Email,
		TaxCode:           billingInfo.TaxCode,
		Phone:             billingInfo.Phone,
		Type:              billingInfo.Type,
		Note:              billingInfo.Note,
		ExpectedIssueDate: billingInfo.ExpectedIssueDate,
	}
}

func (dt *DtoTransformer) ToShipments(shipments []model.Shipment) []dto.Shipment {
	data := make([]dto.Shipment, len(shipments))

	for i, shipment := range shipments {
		shipmentInfo := shipment.ShipmentInfo

		data[i] = dto.Shipment{
			ShipmentID:                  shipmentInfo.ShipmentID,
			ShipmentCode:                shipmentInfo.ShipmentCode,
			SellerID:                    shipmentInfo.SellerID,
			Status:                      shipmentInfo.Status,
			Items:                       dt.ToShipmentItems(shipmentInfo.Items),
			DocumentIds:                 shipmentInfo.DocumentIds,
			FulfillmentChannelID:        shipmentInfo.FulfillmentChannelID,
			FulfillmentChannelName:      shipmentInfo.FulfillmentChannelName,
			IsAuto:                      shipmentInfo.IsAuto,
			ReplacedShipmentID:          shipmentInfo.ReplacedShipmentID,
			OutboundRequestID:           shipmentInfo.OutboundRequestID,
			OutboundRequestAt:           shipmentInfo.OutboundRequestAt,
			WarehouseProviderExportID:   shipmentInfo.WarehouseProviderExportID,
			WarehouseExportAt:           shipmentInfo.WarehouseExportAt,
			WarehouseProviderID:         shipmentInfo.WarehouseProviderID,
			NumberOfPackages:            shipmentInfo.NumberOfPackages,
			Height:                      shipmentInfo.Height,
			Length:                      shipmentInfo.Length,
			Width:                       shipmentInfo.Width,
			TotalWeight:                 shipmentInfo.TotalWeight,
			TransportProviderID:         shipmentInfo.TransportProviderID,
			TransportProviderName:       shipmentInfo.TransportProviderName,
			TransportProviderType:       shipmentInfo.TransportProviderType,
			TransportTrackingID:         shipmentInfo.TransportTrackingID,
			AllowTrialOnDelivery:        shipmentInfo.AllowTrialOnDelivery,
			DeliveryMethodCode:          shipmentInfo.DeliveryMethodCode,
			DeliveryRequestCreatedAt:    shipmentInfo.DeliveryRequestCreatedAt,
			DeliveryRequestCreatorIamID: shipmentInfo.DeliveryRequestCreatorIamID,
			DeliveryRequestID:           shipmentInfo.DeliveryRequestID,
			DeliveryRequestStatus:       shipmentInfo.DeliveryRequestStatus,
			DeliveryTypeGroupID:         shipmentInfo.DeliveryTypeGroupID,
			DeliveryTypeID:              shipmentInfo.DeliveryTypeID,
			Description:                 shipmentInfo.Description,
			ExpectDeliveryDueFrom:       shipmentInfo.ExpectDeliveryDueFrom,
			ExpectDeliveryDueTo:         shipmentInfo.ExpectDeliveryDueTo,
		}
	}
	return data
}

func (dt *DtoTransformer) ToShipmentItems(items []model.ShipmentItem) []dto.ShipmentItem {
	data := make([]dto.ShipmentItem, len(items))

	for i, item := range items {
		data[i] = dto.ShipmentItem{
			Sku:          item.Sku,
			SiteID:       item.SiteID,
			LineItemID:   item.LineItemID,
			Name:         item.Name,
			Services:     item.Services,
			WarehouseID:  item.WarehouseID,
			CancelledQty: item.CancelledQty,
			DeliveredQty: item.DeliveredQty,
			ExportedQty:  item.ExportedQty,
			PackedQty:    item.PackedQty,
			RequestQty:   item.RequestQty,
			HoldID:       item.HoldID,
			HoldItemID:   item.HoldItemID,
			Priority:     item.Priority,
		}
	}
	return data
}

func (dt *DtoTransformer) ToReturns(returnInfo model.ReturnInfo) []dto.ReturnInfo {
	data := make([]dto.ReturnInfo, len(returnInfo))

	for i, rr := range returnInfo {
		data[i] = dto.ReturnInfo{
			ReturnRequestID:           rr.ReturnRequestID,
			PlatformID:                rr.PlatformID,
			ReasonID:                  rr.ReasonID,
			Reason:                    rr.Reason,
			SiteID:                    rr.SiteID,
			WarehouseProviderExportID: rr.WarehouseProviderExportID,
			Status:                    rr.Status,
			Items:                     dt.ToReturnItems(rr.Items),
			CustomerNotify:            rr.CustomerNotify,
			CustomerInstruction:       rr.CustomerInstruction,
			Comment:                   rr.Comment,
			CreatedByEmail:            rr.CreatedByEmail,
			ApprovedByEmail:           rr.ApprovedByEmail,
			CreatedAt:                 rr.CreatedAt,
			UpdatedAt:                 rr.UpdatedAt,
		}
	}

	return data
}

func (dt *DtoTransformer) ToReturnItems(returnItems []model.ReturnRequestItem) []dto.ReturnInfoItem {
	data := make([]dto.ReturnInfoItem, len(returnItems))

	for i, rr := range returnItems {
		data[i] = dto.ReturnInfoItem{
			ID:                    rr.ID,
			SellerID:              rr.SellerID,
			ShipmentID:            rr.ShipmentID,
			Sku:                   rr.Sku,
			ServiceID:             rr.ServiceID,
			Name:                  rr.Name,
			UnitPrice:             rr.UnitPrice,
			Quantity:              rr.Quantity,
			RequestReturnQuantity: rr.RequestReturnQuantity,
			ReturnedQuantity:      rr.ReturnedQuantity,
			ReturnItemStatus:      rr.ReturnItemStatus,
		}
	}
	return data
}

func (dt *DtoTransformer) ToCreator(creator model.OrderCreator) dto.OrderCreator {
	return dto.OrderCreator{
		ID:        creator.ID,
		ProfileID: creator.ProfileID,
		Name:      creator.Name,
		Email:     creator.Email,
		Phone:     creator.Phone,
	}
}

func (dt *DtoTransformer) ToConsultant(consultant model.OrderConsultant) dto.OrderConsultant {
	return dto.OrderConsultant{
		ID:        consultant.ID,
		ProfileID: consultant.ProfileID,
		Name:      consultant.Name,
		Email:     consultant.Email,
		Phone:     consultant.Phone,
		CustomID:  consultant.CustomID,
	}
}

func (dt *DtoTransformer) ToCustomer(customerInfo model.OrderCustomerInfo) dto.OrderCustomer {
	return dto.OrderCustomer{
		ID:        customerInfo.ID,
		ProfileID: customerInfo.ProfileID,
		Name:      customerInfo.Name,
		Email:     customerInfo.Email,
		Phone:     customerInfo.Phone,
	}
}

func (dt *DtoTransformer) ToShippingInfo(deliveryInfo model.OrderDeliveryInfo) dto.OrderShippingInfo {
	return dto.OrderShippingInfo{
		Name:                          deliveryInfo.Name,
		Phone:                         deliveryInfo.Phone,
		Email:                         deliveryInfo.Email,
		AddressID:                     deliveryInfo.AddressID,
		Address:                       deliveryInfo.Address,
		WardID:                        deliveryInfo.WardID,
		WardName:                      deliveryInfo.WardName,
		DistrictID:                    deliveryInfo.DistrictID,
		DistrictName:                  deliveryInfo.DistrictName,
		ProvinceID:                    deliveryInfo.ProvinceID,
		ProvinceName:                  deliveryInfo.ProvinceName,
		FullAddress:                   deliveryInfo.FullAddress,
		Country:                       deliveryInfo.Country,
		Latitude:                      deliveryInfo.Latitude,
		Longitude:                     deliveryInfo.Longitude,
		Note:                          deliveryInfo.Note,
		SiteID:                        deliveryInfo.SiteID,
		CollectionHubPartnerID:        deliveryInfo.CollectionHubPartnerID,
		EstimatedDeliveryDate:         deliveryInfo.EstimatedDeliveryDate,
		ScheduledDeliveryDate:         deliveryInfo.ScheduledDeliveryDate,
		ScheduledDeliveryTimeSlotFrom: deliveryInfo.ScheduledDeliveryTimeSlotFrom,
		ScheduledDeliveryTimeSlotTo:   deliveryInfo.ScheduledDeliveryTimeSlotTo,
	}
}
