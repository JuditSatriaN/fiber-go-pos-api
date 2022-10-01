package sales

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const queryInsertSalesHead = `
	INSERT INTO sales_head (invoice, user_id, total_item, total_price, total_purchase, 
	                        total_tax, total_discount, total_pay)
	VALUES (:invoice, :user_id, :total_item, :total_price, :total_purchase, 
	        :total_tax, :total_discount, :total_pay)
`

func InsertSalesHead(ctx *fiber.Ctx, tx *sqlx.Tx, salesHead model.SalesHead) error {
	_, err := tx.NamedExecContext(ctx.Context(), queryInsertSalesHead, salesHead)
	return err
}
