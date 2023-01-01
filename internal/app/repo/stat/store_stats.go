package stat

import (
	"database/sql"

	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
)

const queryGetStoreStats = `
	SELECT store_id, total_product
	FROM store_stats
	WHERE store_id = $1
`

func GetStoreStats(ctx *fiber.Ctx, storeID string) (model.StoreStats, error) {
	var storeStats model.StoreStats
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &storeStats, queryGetStoreStats, storeID); err != nil {
		if err == sql.ErrNoRows {
			return storeStats, nil
		}

		return storeStats, err
	}

	return storeStats, nil
}

const queryGetTotalProduct = `
	SELECT total_product
	FROM store_stats
	WHERE store_id = $1
`

func GetTotalProduct(ctx *fiber.Ctx, storeID int64) (int64, error) {
	var totalProduct int64
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &totalProduct, queryGetTotalProduct, storeID); err != nil {
		if err == sql.ErrNoRows {
			return totalProduct, nil
		}

		return totalProduct, err
	}

	return totalProduct, nil
}

const queryGetTotalInventory = `
	SELECT total_inventory
	FROM store_stats
	WHERE store_id = $1
`

func GetTotalInventory(ctx *fiber.Ctx, storeID string) (int64, error) {
	var totalInventory int64
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &totalInventory, queryGetTotalInventory, storeID); err != nil {
		if err == sql.ErrNoRows {
			return totalInventory, nil
		}

		return totalInventory, err
	}

	return totalInventory, nil
}

const updateTotalProduct = `
	UPDATE store_stats 
	SET
		total_product = total_product + :total_product,
		update_time = NOW()
	WHERE store_id = :store_id
`

func UpdateTotalProduct(tx *sqlx.Tx, store model.StoreStats) error {
	rows, err := tx.NamedQuery(updateTotalProduct, store)
	if err != nil {
		return err
	}

	defer rows.Close()
	return err
}

const updateTotalInventory = `
	UPDATE store_stats 
	SET
		total_inventory = total_inventory + :total_inventory,
		update_time = NOW()
	WHERE store_id = :store_id
`

func UpdateTotalInventory(tx *sqlx.Tx, store model.StoreStats) error {
	rows, err := tx.NamedQuery(updateTotalInventory, store)
	if err != nil {
		return err
	}

	defer rows.Close()
	return err
}
