ALTER TABLE orders ADD COLUMN site_ids int ARRAY;

CREATE INDEX orders_site_ids_idx ON orders USING GIN (site_ids);