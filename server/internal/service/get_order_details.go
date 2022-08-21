package service

import (
	"net/http"

	httphelper "github.com/dokyan1989/postgresdemo/helper/http"
	jsonhelper "github.com/dokyan1989/postgresdemo/helper/json"
	"github.com/dokyan1989/postgresdemo/server/dto"
	"github.com/dokyan1989/postgresdemo/server/internal/transformer"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
)

func (s *Service) GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	data, err := s.getOrderDetails(w, r)
	if err != nil {
		httphelper.WriteJSON(w, dto.GetOrderDetailsResponse{Code: 1, Message: err.Error()}, http.StatusBadRequest)
		return
	}

	response := dto.GetOrderDetailsResponse{Code: 0, Message: "success", Data: data}
	httphelper.WriteJSON(w, response, http.StatusOK)
}

func (s *Service) getOrderDetails(w http.ResponseWriter, r *http.Request) (*dto.FullSaleOrder, error) {
	var request dto.GetOrderDetailsRequest
	err := s.decodeGetOrderDetailsRequest(&request, r)
	if err != nil {
		s.logger.Error("Error decoding request", zap.Error(err))
		return nil, err
	}

	logger := s.logger.Named("GetOrderDetails").With(zap.Any("request", jsonhelper.Format(request)))

	order, err := s.repo.FindOrderByID(r.Context(), request.OrderID)
	if err != nil {
		logger.Error("Error finding order", zap.Error(err))
		return nil, err
	}

	orderRawData, err := s.repo.FindOrderRawDataByID(r.Context(), request.OrderID)
	if err != nil {
		logger.Error("Error finding order raw data", zap.Error(err))
		return nil, err
	}

	shipments, err := s.repo.ListShipmentsByOrderID(r.Context(), request.OrderID)
	if err != nil {
		logger.Error("Error finding shipments", zap.Error(err))
		return nil, err
	}

	var dt transformer.DtoTransformer
	return dt.ToFullSaleOrder(&order, &orderRawData, shipments), nil
}

func (s *Service) decodeGetOrderDetailsRequest(request *dto.GetOrderDetailsRequest, r *http.Request) error {
	// decode url query
	err := schema.NewDecoder().Decode(request, r.URL.Query())
	if err != nil {
		return err
	}

	// get order id in url param
	request.OrderID = chi.URLParam(r, "orderId")
	return nil
}
