package response

import (
	"reflect"

	"github.com/fiber-go-pos-api/internal/app/constant"
	"github.com/gofiber/fiber/v2"
)

func BuildMetaDataResponse[T constant.IntegerNumber](page T, limit T, total T, links any) constant.MetadataResponse {
	if page == 0 {
		page = constant.DefaultPage
	}

	if limit == 0 {
		limit = constant.DefaultLimit
	}

	pageCount := total / limit
	if (int64(total) % int64(limit)) > 0 {
		pageCount++
	}

	return constant.MetadataResponse{
		Page:       int(page),
		PerPage:    int(limit),
		PageCount:  int64(pageCount),
		TotalCount: int64(total),
		Links:      links,
	}
}

// BuildStandardResponse global function to response in this project
func BuildStandardResponse(ctx *fiber.Ctx, res constant.StandardResponse) error {
	ctx.Set("CONTENT-TYPE", "application/json; charset=utf-8")
	return ctx.Status(res.ResponseCode).JSON(constant.StandardResponse{
		ResponseCode: res.ResponseCode,
		Message:      res.Message,
		Data:         NilSliceToEmptySlice(res.Data),
		Metadata:     res.Metadata,
	})
}

// BuildJSONRes global function to build json response
func BuildJSONRes(ctx *fiber.Ctx, response any) error {
	ctx.Set("CONTENT-TYPE", "application/json; charset=utf-8")
	return ctx.JSON(response)
}

// BuildDatatableRes global function to build datatable response
func BuildDatatableRes(ctx *fiber.Ctx, total int64, data any) error {
	ctx.Set("CONTENT-TYPE", "application/json; charset=utf-8")
	return ctx.JSON(map[string]interface{}{
		"response_code": fiber.StatusOK,
		"total":         total,
		"rows":          NilSliceToEmptySlice(data),
	})
}

// NilSliceToEmptySlice recursively sets nil slices to empty slices
func NilSliceToEmptySlice(inter any) any {
	// original input that can't be modified
	val := reflect.ValueOf(inter)

	switch val.Kind() {
	case reflect.Slice:
		newSlice := reflect.MakeSlice(val.Type(), 0, val.Len())
		if !val.IsZero() {
			// iterate over each element in slice
			for j := 0; j < val.Len(); j++ {
				item := val.Index(j)

				var newItem reflect.Value
				switch item.Kind() {
				case reflect.Struct:
					// recursively handle nested struct
					newItem = reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(item.Interface())))
				default:
					newItem = item
				}

				newSlice = reflect.Append(newSlice, newItem)
			}

		}
		return newSlice.Interface()
	case reflect.Struct:
		// new struct that will be returned
		newStruct := reflect.New(reflect.TypeOf(inter))
		newVal := newStruct.Elem()
		// iterate over input's fields
		for i := 0; i < val.NumField(); i++ {
			newValField := newVal.Field(i)
			valField := val.Field(i)
			switch valField.Kind() {
			case reflect.Slice:
				// recursively handle nested slice
				newValField.Set(reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(valField.Interface()))))
			case reflect.Struct:
				// recursively handle nested struct
				newValField.Set(reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(valField.Interface()))))
			default:
				newValField.Set(valField)
			}
		}

		return newStruct.Interface()
	case reflect.Map:
		// new map to be returned
		newMap := reflect.MakeMap(reflect.TypeOf(inter))
		// iterate over every key value pair in input map
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			// recursively handle nested value
			newV := reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(v.Interface())))
			newMap.SetMapIndex(k, newV)
		}
		return newMap.Interface()
	case reflect.Ptr:
		// dereference pointer
		return NilSliceToEmptySlice(val.Elem().Interface())
	default:
		return inter
	}
}
