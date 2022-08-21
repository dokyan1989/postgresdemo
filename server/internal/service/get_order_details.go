package service

import (
	"net/http"

	httphelper "github.com/dokyan1989/postgresdemo/helper/http"
	jsonhelper "github.com/dokyan1989/postgresdemo/helper/json"
	"github.com/dokyan1989/postgresdemo/model"
	"github.com/dokyan1989/postgresdemo/server/dto"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
)

func (s *Service) GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	var request dto.GetOrderDetailsRequest
	err := schema.NewDecoder().Decode(&request, r.URL.Query())
	if err != nil {
		s.logger.Error("Error decoding request", zap.Error(err))
		httphelper.WriteJSON(w, dto.GetOrderDetailsResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}

	request.OrderID = chi.URLParam(r, "orderId")

	logger := s.logger.Named("GetOrderDetails").With(zap.Any("request", jsonhelper.Format(request)))

	session := s.db.NewSession(nil)

	var order model.Order
	err = session.Select("*").From("orders").Where("id = ?", request.OrderID).LoadOne(&order)
	if err != nil {
		logger.Error("Error finding order", zap.Error(err))
		httphelper.WriteJSON(w, dto.GetOrderDetailsResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}

	var orderRawData model.OrderRawData
	err = session.Select("*").From("orders_raw_data").Where("id = ?", order.ID).LoadOne(&orderRawData)
	if err != nil {
		logger.Error("Error finding order raw data", zap.Error(err))
		httphelper.WriteJSON(w, dto.GetOrderDetailsResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response := dto.GetOrderDetailsResponse{
		Code:    0,
		Message: "success",
		Data: dto.GetOrderDetailsResponseData{
			OrderID:       order.ID,
			PaymentStatus: order.PaymentStatus,
			TotalPaid:     orderRawData.OrderInfo.TotalPaid,
			TerminalCode:  orderRawData.OrderInfo.TerminalCode,
		},
	}

	logger.Info("Get order details successfully", zap.Any("response", jsonhelper.Format(response)))
	httphelper.WriteJSON(w, response, http.StatusOK)
}
