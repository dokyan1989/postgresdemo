-- orders definition

CREATE TABLE IF NOT EXISTS orders (
	id varchar(36) NOT NULL,
	fulfillment_status varchar(50) NOT NULL DEFAULT '',
	payment_status varchar(50) NOT NULL DEFAULT '',
	confirmation_status varchar(50) NOT NULL DEFAULT '',
	customer_phone varchar(32) NOT NULL DEFAULT '',
	customer_name varchar(255) NOT NULL DEFAULT '',
	customer_email varchar(225) NOT NULL DEFAULT '',
	shipping_info_phone varchar(32) NOT NULL DEFAULT '',
	terminal_id int4 NOT NULL DEFAULT 0,
	platform_id int4 NOT NULL DEFAULT 0,
	creator_id varchar(36) NOT NULL,
	consultant_id varchar(36) NOT NULL DEFAULT '',
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT orders_primary PRIMARY KEY (id)
);

CREATE INDEX orders_fulfillment_status_created_at_asc_idx ON orders USING btree (fulfillment_status, created_at ASC);
CREATE INDEX orders_fulfillment_status_updated_at_desc_idx ON orders USING btree (fulfillment_status, updated_at DESC);
CREATE INDEX orders_payment_status_created_at_asc_idx ON orders USING btree (payment_status, created_at ASC);
CREATE INDEX orders_payment_status_updated_at_desc_idx ON orders USING btree (payment_status, updated_at DESC);
CREATE INDEX orders_confirmation_status_created_at_asc_idx ON orders USING btree (confirmation_status, created_at ASC);
CREATE INDEX orders_confirmation_status_updated_at_desc_idx ON orders USING btree (confirmation_status, updated_at DESC);
CREATE INDEX orders_customer_phone_idx ON orders USING btree (customer_phone);
CREATE INDEX orders_customer_name_idx ON orders USING btree (customer_name);
CREATE INDEX orders_customer_email_idx ON orders USING btree (customer_email);
CREATE INDEX orders_shipping_info_phone_idx ON orders USING btree (shipping_info_phone);
CREATE INDEX orders_terminal_id_idx ON orders USING btree (terminal_id);
CREATE INDEX orders_platform_id_idx ON orders USING btree (platform_id);
CREATE INDEX orders_creator_id_idx ON orders USING btree (creator_id);
CREATE INDEX orders_consultant_id_idx ON orders USING btree (consultant_id);

CREATE TABLE IF NOT EXISTS orders_raw_data (
	id varchar(36) NOT NULL,
	order_info jsonb NOT NULL DEFAULT '{}',
	return_info jsonb NOT NULL DEFAULT '[]',
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT orders_raw_data_primary PRIMARY KEY (id)
);

-- shipments definition

CREATE TABLE IF NOT EXISTS shipments (
	id varchar(36) NOT NULL,
	order_id varchar(36) NOT NULL,
	seller_id int4 NOT NULL DEFAULT 0,
	site_id int4 NOT NULL DEFAULT 0,
	status varchar(255) NOT NULL,
	shipment_info jsonb NOT NULL DEFAULT '{}',
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT shipments_primary PRIMARY KEY (id, order_id)
);
CREATE INDEX shipments_order_id_idx ON shipments USING btree (order_id);
CREATE INDEX shipments_status_created_at_asc_idx ON shipments USING btree (status, created_at ASC);
CREATE INDEX shipments_status_updated_at_desc_idx ON shipments USING btree (status, updated_at DESC);
CREATE INDEX shipments_seller_id_created_at_asc_idx ON shipments USING btree (seller_id, created_at ASC);
CREATE INDEX shipments_seller_id_updated_at_desc_idx ON shipments USING btree (seller_id, updated_at DESC);
CREATE INDEX shipments_site_id_created_at_asc_idx ON shipments USING btree (site_id, created_at ASC);
CREATE INDEX shipments_site_id_updated_at_desc_idx ON shipments USING btree (site_id, updated_at DESC);


-- invoices definition

CREATE TABLE IF NOT EXISTS invoices (
	order_id varchar(36) NOT NULL,
	shipment_id varchar(36) NOT NULL,
	seller_id int4 NOT NULL,
	invoice_info jsonb NOT NULL DEFAULT '{}',
	created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT invoices_primary PRIMARY KEY (order_id, shipment_id, seller_id)
);
CREATE INDEX invoices_order_id_idx ON invoices USING btree (order_id);
CREATE INDEX invoices_shipment_id_seller_id_idx ON invoices USING btree (shipment_id, seller_id);
CREATE INDEX invoices_seller_id_created_at_asc_idx ON invoices USING btree (seller_id, created_at ASC);
CREATE INDEX invoices_seller_id_updated_at_desc_idx ON invoices USING btree (seller_id, updated_at DESC);