package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	httphelper "github.com/dokyan1989/postgresdemo/helper/http"
	jsonhelper "github.com/dokyan1989/postgresdemo/helper/json"
	"github.com/dokyan1989/postgresdemo/model"
	"github.com/dokyan1989/postgresdemo/server/dto"
	"github.com/dokyan1989/postgresdemo/server/internal/repository"
	"github.com/dokyan1989/postgresdemo/server/internal/transformer"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (s *Service) ListOrdersByPlatform(w http.ResponseWriter, r *http.Request) {
	data, err := s.listOrdersByPlatform(w, r)
	if err != nil {
		httphelper.WriteJSON(w, dto.ListOrdersByPlatformResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response := dto.ListOrdersByPlatformResponse{Code: 0, Message: "success", Data: data}
	httphelper.WriteJSON(w, response, http.StatusOK)
}

func (s *Service) listOrdersByPlatform(w http.ResponseWriter, r *http.Request) (dto.ListOrdersResponseData, error) {
	var data dto.ListOrdersResponseData

	var request dto.ListOrdersByPlatformRequest
	err := s.decodeListOrdersByPlatformRequest(&request, r)
	if err != nil {
		s.logger.Error("Error decoding request", zap.Error(err))
		return data, err
	}

	logger := s.logger.Named("ListOrdersByPlatform").With(zap.Any("request", jsonhelper.Format(request)))

	orders, err := s.repo.ListOrders(r.Context(), repository.ListOrdersOptions{
		PlatformID:         request.PlatformID,
		SellerID:           request.SellerID,
		TerminalID:         request.TerminalID,
		SiteID:             request.SiteID,
		CreatorID:          request.CreatorID,
		ConsultantID:       request.ConsultantID,
		HoldStatus:         request.HoldStatus,
		FulfillmentStatus:  request.FulfillmentStatus,
		PaymentStatus:      request.PaymentStatus,
		ConfirmationStatus: request.ConfirmationStatus,
		Customer:           request.Customer,
		OrderID:            request.OrderID,
		CreatedAtGte:       request.CreatedAtGte,
		CreatedAtLte:       request.CreatedAtLte,
		SortBy:             request.SortBy,
		SortOrder:          request.SortOrder,
		Limit:              request.Limit,
		Offset:             request.Offset,
	})
	if err != nil {
		logger.Error("Error finding orders", zap.Error(err))
		return data, err
	}

	ordersRawData, err := s.repo.ListOrdersRawDataByOrderIDs(r.Context(), model.OrderList(orders).GetIDs())
	if err != nil {
		logger.Error("Error finding orders raw data", zap.Error(err))
		return data, err
	}

	var dt transformer.DtoTransformer
	data.Orders = dt.ToBaseSaleOrders(orders, ordersRawData)

	return data, nil
}

func (s *Service) decodeListOrdersByPlatformRequest(request *dto.ListOrdersByPlatformRequest, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.Error("Error decoding request", zap.Error(err))
		return err
	}

	platformId, err := strconv.ParseInt(chi.URLParam(r, "platformId"), 10, 64)
	if err != nil {
		s.logger.Error("Error parsing platform id", zap.Error(err))
		return err
	}

	if err := request.Validate(); err != nil {
		s.logger.Error("Error validating request", zap.Error(err))
		return err
	}

	request.PlatformID = platformId
	return nil
}
