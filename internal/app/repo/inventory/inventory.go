package inventory

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
	errorPkg "github.com/fiber-go-pos-api/internal/pkg/error"
)

const queryGetALlInventory = `
	SELECT i.id, p.plu, p.name, u.id AS unit_id, u.name AS unit_name, p.barcode, p.ppn, 
	       i.multiplier, i.stock, i.price, i.member_price, i.purchase, i.discount, 
	       i.discount_percentage 
	FROM inventories i
	INNER JOIN products p ON i.plu = p.plu
	INNER JOIN units u ON i.unit_id = u.id
	WHERE $1 = '' OR p.value_text_search @@ plainto_tsquery($1)
	LIMIT $2
	OFFSET $3
`

func GetALlInventory(ctx *fiber.Ctx, search string, limit int, offset int) ([]model.Inventory, error) {
	var inventory []model.Inventory
	db := postgresPkg.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &inventory, queryGetALlInventory, search, limit, offset); err != nil {
		return inventory, err
	}

	return inventory, nil
}

const queryGetTotalDataInventory = `
	SELECT COUNT(*) OVER (ROWS BETWEEN CURRENT ROW AND 1000 FOLLOWING) AS total_count
	FROM inventories i
	INNER JOIN products p on i.plu = p.plu
	WHERE $1 = '' OR p.value_text_search @@ plainto_tsquery($1)
	LIMIT 1
`

func GetTotalDataInventory(ctx *fiber.Ctx, search string) (int64, error) {
	var totalData int64
	db := postgresPkg.GetPgConn()

	if err := db.GetContext(ctx.Context(), &totalData, queryGetTotalDataInventory, search); err != nil {
		if errorPkg.IsErrNoRows(err) {
			return totalData, nil
		}
		return totalData, err
	}

	return totalData, nil
}

const queryGetInventoryByID = `
	SELECT id, plu
	FROM inventories
	WHERE id = $1
`

func GetInventoryByID(ctx *fiber.Ctx, ID int64) (model.Inventory, bool, error) {
	var inventory model.Inventory
	db := postgresPkg.GetPgConn()
	if err := db.GetContext(ctx.Context(), &inventory, queryGetInventoryByID, ID); err != nil {
		if errorPkg.IsErrNoRows(err) {
			return inventory, false, nil
		}

		return inventory, false, err
	}

	return inventory, true, nil
}

const queryGetInventoryByPLU = `
	SELECT i.id, p.plu, p.name, u.id AS unit_id, u.name AS unit_name, p.barcode, p.ppn, 
	       i.multiplier, i.stock, i.price, i.member_price, i.purchase, i.discount, i.discount_percentage
	FROM inventories i
	INNER JOIN products p on i.plu = p.plu
	INNER JOIN units u on i.unit_id = u.id
	WHERE i.plu = ($1)
`

func GetInventoryByPLU(ctx *fiber.Ctx, ID int64) ([]model.Inventory, error) {
	var inventories []model.Inventory
	db := postgresPkg.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &inventories, queryGetInventoryByPLU, ID); err != nil {
		return inventories, err
	}

	return inventories, nil
}

const queryInsertInventory = `
	INSERT INTO inventories (plu, unit_id, multiplier, stock, price, member_price, purchase, discount, discount_percentage)
	VALUES (:plu, :unit_id, :multiplier, :stock, :price, :member_price, :purchase, :discount, :discount_percentage)
`

func InsertInventory(tx *sqlx.Tx, inventory model.Inventory) error {
	rows, err := tx.NamedQuery(queryInsertInventory, inventory)
	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}

const queryDeleteInventory = `
	DELETE FROM inventories
	WHERE id = $1
`

func DeleteInventory(tx *sqlx.Tx, ID int64) error {
	_, err := tx.Exec(queryDeleteInventory, ID)
	return err
}

const queryUpdateInventory = `
	UPDATE inventories 
	SET multiplier = :multiplier,
		stock = :stock,
		price = :price,
		member_price = :member_price,
		purchase = :purchase,
		discount = :discount,
		discount_percentage = :discount_percentage,
	    update_time = NOW()
	WHERE id = :id
`

func UpdateInventory(ctx *fiber.Ctx, inventory model.Inventory) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryUpdateInventory, inventory)
	return err
}

const queryUpdateStockInventory = `
	UPDATE inventories 
	SET stock = :stock,
	    update_time = NOW()
	WHERE id = :id
`

func UpdateStockInventory(ctx *fiber.Ctx, inventory model.Inventory) error {
	db := postgresPkg.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), queryUpdateStockInventory, inventory)
	return err
}

const queryUpdateStockAfterSales = `
	UPDATE inventories 
	SET stock = stock - :qty,
	    update_time = NOW()
	WHERE id = :id
`

func UpdateStockAfterSales(tx *sqlx.Tx, data model.UpdateStockAfterSalesData) error {
	rows, err := tx.NamedQuery(queryUpdateStockAfterSales, data)
	if err != nil {
		return err
	}

	defer rows.Close()
	return nil
}
