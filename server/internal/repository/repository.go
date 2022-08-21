package repository

import (
	"context"

	"github.com/dokyan1989/postgresdemo/model"
	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type Repository interface {
	OrderInterface
	ShipmentInterface
}

type repositoryImpl struct {
	db     *dbr.Connection
	logger *zap.Logger
}

func New(db *dbr.Connection, logger *zap.Logger) Repository {
	return &repositoryImpl{db: db, logger: logger}
}

type OrderInterface interface {
	FindOrderByID(ctx context.Context, orderID string) (model.Order, error)
	FindOrderRawDataByID(ctx context.Context, orderID string) (model.OrderRawData, error)
	ListOrders(ctx context.Context, opts ListOrdersOptions) ([]model.Order, error)
	ListOrdersRawDataByOrderIDs(ctx context.Context, orderIDs []string) ([]model.OrderRawData, error)
}

type ShipmentInterface interface {
	ListShipmentsByOrderID(ctx context.Context, orderID string) ([]model.Shipment, error)
}
