package repository

import (
	"context"
	"fmt"
	"strings"

	dbrhelper "github.com/dokyan1989/postgresdemo/helper/dbr"
	"github.com/dokyan1989/postgresdemo/model"
	"github.com/gocraft/dbr/v2"
)

func (r *repositoryImpl) FindOrderByID(ctx context.Context, orderID string) (model.Order, error) {
	session := r.db.NewSession(nil)

	stmt := session.Select("*").From("orders").Where("id = ?", orderID)
	r.logger.Info(fmt.Sprintf("%s - %s", "[FindOrderByIDQuery]", dbrhelper.SprintSelect(stmt)))

	var order model.Order
	err := stmt.LoadOne(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repositoryImpl) FindOrderRawDataByID(ctx context.Context, orderID string) (model.OrderRawData, error) {
	session := r.db.NewSession(nil)

	stmt := session.Select("*").From("orders_raw_data").Where("id = ?", orderID)
	r.logger.Info(fmt.Sprintf("%s - %s", "[FindOrderRawDataByIDQuery]", dbrhelper.SprintSelect(stmt)))

	var orderRawData model.OrderRawData
	err := stmt.LoadOne(&orderRawData)
	if err != nil {
		return orderRawData, err
	}

	return orderRawData, nil
}

func (r *repositoryImpl) ListOrders(ctx context.Context, opts ListOrdersOptions) ([]model.Order, error) {
	session := r.db.NewSession(nil)

	stmt := r.buildListOrdersStmt(opts, session)
	r.logger.Info(fmt.Sprintf("%s - %s", "[ListOrdersQuery]", dbrhelper.SprintSelect(stmt)))

	var orders []model.Order
	_, err := stmt.Load(&orders)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (r *repositoryImpl) ListOrdersRawDataByOrderIDs(ctx context.Context, orderIDs []string) ([]model.OrderRawData, error) {
	session := r.db.NewSession(nil)

	stmt := session.Select("*").From("orders_raw_data").Where("id IN ?", orderIDs)
	r.logger.Info(fmt.Sprintf("%s - %s", "[ListOrdersRawDataByOrderIDsQuery]", dbrhelper.SprintSelect(stmt)))

	var ordersRawData []model.OrderRawData
	_, err := stmt.Load(&ordersRawData)
	if err != nil {
		return ordersRawData, err
	}

	return ordersRawData, nil
}

func (r *repositoryImpl) buildListOrdersStmt(opts ListOrdersOptions, session *dbr.Session) *dbr.SelectStmt {
	stmt := session.Select("*").From("orders")

	if opts.PlatformID != 0 {
		stmt = stmt.Where("platform_id = ?", opts.PlatformID)
	}

	if opts.TerminalID != 0 {
		stmt = stmt.Where("terminal_id = ?", opts.TerminalID)
	}

	if opts.CreatorID != "" {
		stmt = stmt.Where("creator_id = ?", opts.CreatorID)
	}

	if opts.ConsultantID != "" {
		stmt = stmt.Where("consultant_id = ?", opts.ConsultantID)
	}

	if opts.HoldStatus.Valid {
		stmt = stmt.Where("hold_status = ?", opts.HoldStatus.Bool)
	}

	if opts.FulfillmentStatus != "" {
		stmt = stmt.Where("fulfillment_status = ?", opts.FulfillmentStatus)
	}

	if opts.PaymentStatus != "" {
		stmt = stmt.Where("payment_status = ?", opts.PaymentStatus)
	}

	if opts.ConfirmationStatus != "" {
		stmt = stmt.Where("confirmation_status = ?", opts.ConfirmationStatus)
	}

	if opts.Customer != "" {
		stmt = stmt.Where("customer_phone = ? OR customer_name = ? OR customer_email = ? OR shipping_info_phone = ?",
			opts.Customer, opts.Customer, opts.Customer, opts.Customer)
	}

	if opts.OrderID != "" {
		stmt = stmt.Where("order_id = ?", opts.OrderID)
	}

	if opts.SiteID != 0 {
		stmt = stmt.Where("site_ids @> '{?}'", opts.SiteID)
	}

	if opts.CreatedAtGte.Valid {
		stmt = stmt.Where("created_at >= ?", opts.CreatedAtGte.Time)
	}

	if opts.CreatedAtLte.Valid {
		stmt = stmt.Where("created_at <= ?", opts.CreatedAtLte.Time)
	}

	isAsc := false
	if strings.ToLower(opts.SortOrder) == "ASC" {
		isAsc = true
	}

	limit := uint64(10)
	if opts.Limit != 0 {
		limit = opts.Limit
	}

	offset := uint64(0)
	if opts.Offset != 0 {
		offset = opts.Offset
	}

	stmt = stmt.OrderDir(opts.SortBy, isAsc).Limit(limit).Offset(offset)

	return stmt
}
