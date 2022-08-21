package repository

import (
	"context"
	"fmt"

	dbrhelper "github.com/dokyan1989/postgresdemo/helper/dbr"
	"github.com/dokyan1989/postgresdemo/model"
)

func (r *repositoryImpl) ListShipmentsByOrderID(ctx context.Context, orderID string) ([]model.Shipment, error) {
	session := r.db.NewSession(nil)

	stmt := session.Select("*").From("shipments").Where("order_id = ?", orderID)
	r.logger.Info(fmt.Sprintf("%s - %s", "[ListShipmentsByOrderIDQuery]", dbrhelper.SprintSelect(stmt)))

	var shipments []model.Shipment
	_, err := stmt.Load(&shipments)
	if err != nil {
		return shipments, err
	}

	return shipments, nil
}
