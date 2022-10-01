package product

import (
	"fmt"

	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/fiber-go-pos-api/internal/app/model"
	"github.com/gofiber/fiber/v2"

	productRepo "github.com/fiber-go-pos-api/internal/app/repo/product"
	statRepo "github.com/fiber-go-pos-api/internal/app/repo/stat"
	postgresPkg "github.com/fiber-go-pos-api/internal/pkg/database/postgres"
	requestPkg "github.com/fiber-go-pos-api/internal/pkg/request"
)

// GetAllDTProduct : Get List Of Product for Datatable
func GetAllDTProduct(ctx *fiber.Ctx, page int, limit int, search string) (model.ListProductDataResponse, error) {
	offset := requestPkg.BuildOffset(page, limit)

	products, err := productRepo.GetALlProducts(ctx, search, limit, offset)
	if err != nil {
		return model.ListProductDataResponse{}, err
	}

	totalProduct, err := statRepo.GetTotalProduct(ctx, constant.DefaultStoreID)
	if err != nil {
		return model.ListProductDataResponse{}, err
	}

	return model.ListProductDataResponse{
		Total: totalProduct,
		Data:  products,
	}, nil
}

// GetAllProduct : Get List Of Product
func GetAllProduct(ctx *fiber.Ctx, page int, limit int, search string) ([]model.Product, error) {
	offset := requestPkg.BuildOffset(page, limit)

	products, err := productRepo.GetALlProducts(ctx, search, limit, offset)
	if err != nil {
		return []model.Product{}, err
	}

	return products, nil
}

func GetProductByPLU(ctx *fiber.Ctx, ID int64) (model.Product, error) {
	product, found, err := productRepo.GetProductByPLU(ctx, ID)
	if err != nil {
		return model.Product{}, err
	}

	if !found {
		return model.Product{}, fmt.Errorf("product dengan nama : %s tidak ditemukan", product.Name)
	}

	return product, nil
}

func GetProductByPLUOrBarcode(ctx *fiber.Ctx, search string) (model.Product, error) {
	product, found, err := productRepo.GetProductByPLUOrBarcode(ctx, search)
	if err != nil {
		return model.Product{}, err
	}

	if !found {
		return model.Product{}, constant.ErrNoDataFound
	}

	return product, nil
}

func InsertProduct(ctx *fiber.Ctx, product model.Product) error {
	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := productRepo.InsertProduct(tx, product); err != nil {
		return err
	}

	if err := statRepo.UpdateTotalProduct(tx, model.StoreStats{
		StoreID:      constant.DefaultStoreID,
		TotalProduct: 1,
	}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpdateProduct(ctx *fiber.Ctx, product model.Product) error {
	if _, err := GetProductByPLU(ctx, product.ID); err != nil {
		return err
	}
	return productRepo.UpdateProduct(ctx, product)
}

func DeleteProduct(ctx *fiber.Ctx, ID int64) error {
	if _, err := GetProductByPLU(ctx, ID); err != nil {
		return err
	}

	tx, err := postgresPkg.BeginTxx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := productRepo.DeleteProduct(ctx, tx, ID); err != nil {
		return err
	}

	if err := statRepo.UpdateTotalProduct(tx, model.StoreStats{
		StoreID:      constant.DefaultStoreID,
		TotalProduct: -1,
	}); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func UpsertProduct(ctx *fiber.Ctx, product model.Product) error {
	_, found, err := productRepo.GetProductByPLU(ctx, product.ID)
	if err != nil {
		return err
	}

	if !found {
		if err := InsertProduct(ctx, product); err != nil {
			return err
		}
	} else {
		if err := UpdateProduct(ctx, product); err != nil {
			return err
		}
	}

	return nil
}
