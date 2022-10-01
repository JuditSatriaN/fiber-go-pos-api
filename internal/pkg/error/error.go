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

// IsErrNotFound checks whether err is/contains errNotFound
func IsErrNotFound(err error) bool {
	if err == nil {
		return false
	}

	// check sql.ErrNoRows
	return errors.Is(err, sql.ErrNoRows)
}

func ConvertErrorStartsWith(field string, param string) error {
	return fmt.Errorf("Field %s harus dimulai dengan kata %s ", field, param)
}
