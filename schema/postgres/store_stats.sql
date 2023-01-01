CREATE TABLE IF NOT EXISTS store_stats (
    store_id BIGINT NOT NULL PRIMARY KEY,
    total_product BIGINT NOT NULL DEFAULT 0,
    total_inventory BIGINT NOT NULL DEFAULT 0,
    create_time TIMESTAMP NOT NULL DEFAULT NOW(),
    update_time TIMESTAMP
);

INSERT INTO
    store_stats (store_id, total_product, total_inventory)
VALUES
    ('1', 0, 0) ON CONFLICT DO NOTHING;