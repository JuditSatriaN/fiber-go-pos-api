package error

import (
	"database/sql"
	"errors"
	"fmt"
)

// ErrNotFound wraps err as errNotFound
func ErrNotFound(err error) error {
	if err == nil {
		return nil
	}

	//TODO: Add something in this maybe log or etc
	return err
}

// IsErrNoRows checks whether err is/contains ErrNoRows
func IsErrNoRows(err error) bool {
	if err == nil {
		return false
	}

	// check sql.ErrNoRows
	return errors.Is(err, sql.ErrNoRows)
}

func ConvertErrorStartsWith(field string, param string) error {
	return fmt.Errorf("Field %s harus dimulai dengan kata %s ", field, param)
}

func ConvertErrorRequired(field string) error {
	return fmt.Errorf("field %s wajib diisi", field)
}

func ConvertErrorMax(field string, max string) error {
	return fmt.Errorf("field %s melebihi %s huruf", field, max)
}
