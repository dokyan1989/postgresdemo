package service

import (
	"net/http"
	"strconv"

	httphelper "github.com/dokyan1989/postgresdemo/helper/http"
	jsonhelper "github.com/dokyan1989/postgresdemo/helper/json"
	"github.com/dokyan1989/postgresdemo/model"
	"github.com/dokyan1989/postgresdemo/server/dto"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
)

func (s *Service) ListOrdersByPlatform(w http.ResponseWriter, r *http.Request) {
	var request dto.ListOrdersByPlatformRequest
	err := schema.NewDecoder().Decode(&request, r.URL.Query())
	if err != nil {
		s.logger.Error("Error decoding request", zap.Error(err))
		httphelper.WriteJSON(w, dto.ListOrdersByPlatformResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}

	platformId, err := strconv.ParseInt(chi.URLParam(r, "platformId"), 10, 64)
	if err != nil {
		s.logger.Error("Error parsing platform id", zap.Error(err))
		httphelper.WriteJSON(w, dto.ListOrdersByPlatformResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}
	request.PlatformID = int(platformId)

	logger := s.logger.Named("ListOrdersByPlatform").With(zap.Any("request", jsonhelper.Format(request)))

	session := s.db.NewSession(nil)

	var orders model.OrderList
	_, err = session.Select("*").From("orders").Where("platform_id = ?", request.PlatformID).Limit(10).Load(&orders)
	if err != nil {
		logger.Error("Error finding orders", zap.Error(err))
		httphelper.WriteJSON(w, dto.ListOrdersByPlatformResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}

	var orderRawData []model.OrderRawData
	_, err = session.Select("*").From("orders_raw_data").Where("id IN ?", orders.GetIDs()).Load(&orderRawData)
	if err != nil {
		logger.Error("Error finding orders raw data", zap.Error(err))
		httphelper.WriteJSON(w, dto.ListOrdersByPlatformResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response := dto.ListOrdersByPlatformResponse{
		Code:    0,
		Message: "success",
		Data: dto.ListOrdersResponseData{
			Orders: []dto.ListOrdersBaseOrder{},
		},
	}

	logger.Info("Get order details successfully", zap.Any("response", jsonhelper.Format(response)))
	httphelper.WriteJSON(w, response, http.StatusOK)
}
