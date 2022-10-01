package constant

import (
	"errors"
)

// Error Fields Name

const ErrFieldStartsWith = "startswith"
const ErrFieldRequired = "required"
const ErrFieldMax = "max"

// Error Constanta

const ErrInvalidORExpiredJWT = "Invalid or expired JWT"
const ErrMissingOrMalformedJWT = "Missing or malformed JWT"
const ErrUserUnauthorized = "Anda tidak mempunyai akses untuk aplikasi ini"

// Error Message (Will returned to user)

var ErrNoDataFound = errors.New("data tidak ditemukan")
var ErrPageMustBeNumber = errors.New("page must be number")
var ErrLimitMustBeNumber = errors.New("limit must be number")
var ErrShopIDMustBeNumber = errors.New("shop_id must be number")
var ErrInvalidLogin = errors.New("username atau password yang anda masukkan salah")
var ErrUserNotFound = errors.New("user yang anda masukkan tidak ada di dalam database")
