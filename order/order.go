package order

import "constraints"

func Asc[T constraints.Ordered](a, b T) bool {
	return a < b
}

func Desc[T constraints.Ordered](a, b T) bool {
	return a > b
}
