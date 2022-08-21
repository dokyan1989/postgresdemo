package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dokyan1989/postgresdemo/helper/random"
	"github.com/dokyan1989/postgresdemo/server/internal/repository"
	"github.com/dokyan1989/postgresdemo/server/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gocraft/dbr/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"

	modelgen "github.com/dokyan1989/postgresdemo/model/gen"
)

var dsn string = "postgresql://postgres:123456@localhost:5432/testdb?sslmode=disable"
var conn *dbr.Connection
var logger *zap.Logger
var seedNum int

func main() {
	var err error
	conn, err = dbr.Open("postgres", dsn, nil)
	if err != nil {
		log.Fatalf("error opening postgres: %s", err)
	}
	conn.SetMaxOpenConns(10)

	logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatalf("error creating the logger: %s", err)
	}

	app := cli.NewApp()
	app.Name = "postgresdemo"
	app.Description = "Postgres Demo"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:        "seednum",
			Value:       10,
			Usage:       "Number of generated orders",
			Destination: &seedNum,
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "start http server",
			Action: start,
		},
		{
			Name:   "migration",
			Usage:  "data migration",
			Action: migration,
		},
		{
			Name:   "seed",
			Usage:  "data seed",
			Action: seed,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(cliCtx *cli.Context) error {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	repo := repository.New(conn, logger)
	s := service.NewService(repo, logger)

	r.Get("/orders/{orderId}", s.GetOrderDetails)
	r.Post("/orders-by-platform/{platformId}", s.ListOrdersByPlatform)
	// r.Post("/orders-by-seller/{sellerId}", s.ListOrdersBySeller)

	http.ListenAndServe(":3000", r)

	return nil
}

func migration(cliCtx *cli.Context) error {
	sess := conn.NewSession(nil)
	driver, err := postgres.WithInstance(sess.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error init driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://server/migrations", "testdb", driver)
	if err != nil {
		return fmt.Errorf("error init migration: %s", err)
	}

	err = m.Up()
	if err != nil {
		return fmt.Errorf("error running migration: %s", err)
	}

	return nil
}

func seed(cliCtx *cli.Context) error {
	for i := 0; i < 2000; i++ {
		t := time.Now()
		var total int64 = 0

		jobs := make(chan struct{}, seedNum)
		rowCreated := make(chan int64, seedNum)

		// logger.Info("Seed num", zap.Int("seed_num", seedNum))

		for w := 1; w <= 2; w++ {
			go func(w int) {
				for range jobs {
					n, err := doSeed(conn)
					if err != nil {
						logger.Error("doSeed error", zap.Error(err))
						return
					}
					rowCreated <- n
				}
			}(w)
		}

		for orderIndex := 0; orderIndex < seedNum; orderIndex++ {
			jobs <- struct{}{}
		}
		close(jobs)

		var i int = 0
		for orderIndex := 0; orderIndex < seedNum; orderIndex++ {
			n := <-rowCreated
			total += n
			i++
		}

		logger.Info(fmt.Sprintf("Seeding concurrency finished in %f seconds for inserting %d records", time.Since(t).Seconds(), total))
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func doSeed(conn *dbr.Connection) (int64, error) {
	var rowCreatedCount int64 = 0

	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.RollbackUnlessCommitted()

	orderRawData, err := modelgen.GenOrderRawData()
	if err != nil {
		return 0, err
	}

	err = tx.InsertInto("orders_raw_data").
		Columns(
			"id",
			"order_info",
			"return_info",
			"created_at",
			"updated_at",
		).
		Record(orderRawData).
		Returning("id").
		Load(&orderRawData.ID)
	if err != nil {
		return 0, err
	}
	rowCreatedCount++

	// Insert generated order
	order := modelgen.GenOrder(orderRawData)
	err = tx.InsertInto("orders").
		Columns(
			"id",
			"fulfillment_status",
			"payment_status",
			"hold_status",
			"confirmation_status",
			"customer_phone",
			"customer_name",
			"customer_email",
			"shipping_info_phone",
			"terminal_id",
			"platform_id",
			"creator_id",
			"consultant_id",
			"site_ids",
			"created_at",
			"updated_at",
		).
		Record(order).
		Returning("id").
		Load(&order.ID)
	if err != nil {
		return 0, err
	}
	rowCreatedCount++

	// logger.Info("Order is created successfully", zap.String("order_id", order.ID))

	shipmentsNum := random.Int(1, 3)
	for shipmentIndex := 0; shipmentIndex < shipmentsNum; shipmentIndex++ {
		shipment := modelgen.GenShipment(order.ID)
		err := tx.InsertInto("shipments").
			Columns(
				"id",
				"order_id",
				"seller_id",
				"site_id",
				"status",
				"shipment_info",
				"created_at",
				"updated_at",
			).
			Record(shipment).
			Returning("id").
			Load(&shipment.ID)
		if err != nil {
			return 0, err
		}
		rowCreatedCount++
		// logger.Info("Shipment is created successfully", zap.String("shipment_id", shipment.ID))

		invoice := modelgen.GenInvoice(order.ID, shipment.ID, shipment.SellerID)
		_, err = tx.InsertInto("invoices").
			Columns(
				"order_id",
				"shipment_id",
				"seller_id",
				"invoice_info",
				"created_at",
				"updated_at",
			).
			Record(invoice).
			Exec()
		if err != nil {
			return 0, err
		}
		rowCreatedCount++
		// logger.Info("Invoice is created successfully", zap.String("invoice_id", fmt.Sprintf("%s_%s_%d", invoice.OrderID, invoice.ShipmentID, invoice.SellerID)))
	}

	tx.Commit()
	return rowCreatedCount, nil
}
