ALTER TABLE orders ADD COLUMN hold_status boolean;

CREATE INDEX orders_hold_status_created_at_asc_idx ON orders USING btree (hold_status, created_at ASC);
CREATE INDEX orders_hold_status_updated_at_desc_idx ON orders USING btree (hold_status, updated_at DESC);