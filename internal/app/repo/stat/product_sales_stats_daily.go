package stat

import (
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
)

const queryUpsertTotalSold = `
	INSERT INTO product_sales_stats_daily (date_sold, plu, total_sold)
	VALUES (:date_sold, :plu, :total_sold)
	ON CONFLICT (date_sold, plu) DO UPDATE
		SET total_sold  = product_sales_stats_daily.total_sold + EXCLUDED.total_sold,
			update_time = NOW();
`

func BulkUpsertTotalSold(tx *sqlx.Tx, data []model.ProductSalesStatsDaily) error {
	rows, err := tx.NamedQuery(queryUpsertTotalSold, data)
	if err != nil {
		return err
	}
	defer rows.Close()

	return err
}

const queryGetProductSalesStatsDaily = `
	SELECT pd.plu, p.name, pd.total_sold
	FROM product_sales_stats_daily pd
	JOIN products p ON pd.plu = p.plu
	WHERE pd.date_sold = CURRENT_DATE
	ORDER BY pd.total_sold desc
	LIMIT $1;
`

func GetProductSalesStatsDaily(ctx *fiber.Ctx, limit int) ([]model.ProductSalesStatsDaily, error) {
	var productSales []model.ProductSalesStatsDaily
	db := postgresPkg.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &productSales, queryGetProductSalesStatsDaily, limit); err != nil {
		return []model.ProductSalesStatsDaily{}, err
	}

	return productSales, nil
}
