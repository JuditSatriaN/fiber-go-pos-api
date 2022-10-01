CREATE TABLE IF NOT EXISTS products (
    id BIGINT NOT NULL PRIMARY KEY,
    shop_id BIGINT NOT NULL,
    unit_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    stock BIGINT NOT NULL DEFAULT 0,
    barcode VARCHAR(15) NOT NULL DEFAULT '0',
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    member_price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    discount NUMERIC(10, 2) NOT NULL DEFAULT 0,
    discount_percentage SMALLINT NOT NULL DEFAULT 0,
    purchase NUMERIC(10, 2) NOT NULL DEFAULT 0,
    ppn BOOLEAN NOT NULL DEFAULT FALSE,
    create_time TIMESTAMP NOT NULL DEFAULT NOW(),
    update_time TIMESTAMP,
    value_text_search TSVECTOR
);

CREATE INDEX IF NOT EXISTS products_shop_id_barcode_idx ON products (shop_id, barcode);

CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE EXTENSION IF NOT EXISTS btree_gin;

CREATE INDEX IF NOT EXISTS products_shop_id_value_text_idx ON products USING GIN (shop_id, value_text_search);

CREATE OR REPLACE FUNCTION products_upsert_search_trigger() RETURNS trigger AS
$$
BEGIN
    new.value_text_search := TO_TSVECTOR(new.id || ' ' || new.shop_id || ' ' || new.name || ' ' || new.barcode);
    RETURN new;
END
$$ LANGUAGE plpgsql;

BEGIN;

DROP TRIGGER IF EXISTS products_upsert_search ON products;

CREATE TRIGGER products_upsert_search BEFORE
INSERT
    OR
UPDATE
    ON products FOR EACH ROW EXECUTE PROCEDURE products_upsert_search_trigger();

COMMIT;