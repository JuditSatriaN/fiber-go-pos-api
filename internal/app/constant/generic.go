package constant

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type IntegerNumber interface {
	int | int8 | int16 | int32 | int64
}
