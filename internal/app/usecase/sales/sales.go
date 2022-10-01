package sales

import (
	"time"

	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	inventoryRepo "github.com/fiber-go-pos-api/internal/app/repo/inventory"
	salesRepo "github.com/fiber-go-pos-api/internal/app/repo/sales"
	statRepo "github.com/fiber-go-pos-api/internal/app/repo/stat"
	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
)

func InsertSales(ctx *fiber.Ctx, sales model.Sales) error {
	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := salesRepo.InsertSalesHead(ctx, tx, sales.Head); err != nil {
		return err
	}

	if err := salesRepo.InsertSalesDetail(ctx, tx, sales.Detail); err != nil {
		return err
	}

	for _, salesDetailData := range sales.Detail {
		salesData := model.UpdateStockAfterSalesData{
			ID:  salesDetailData.ID,
			Qty: salesDetailData.Qty,
		}
		if err := inventoryRepo.UpdateStockAfterSales(tx, salesData); err != nil {
			return err
		}
	}

	productSalesStats := make([]model.ProductSalesStatsDaily, len(sales.Detail))
	for idx, data := range sales.Detail {
		productSalesStats[idx] = model.ProductSalesStatsDaily{
			DateSold:  time.Now(),
			ID:        data.ID,
			TotalSold: int64(data.Qty),
		}
	}

	if err := statRepo.BulkUpsertTotalSold(tx, productSalesStats); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
