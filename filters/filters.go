package filters

import (
	"constraints"
	"reflect"
	"strings"
)

var (
	nillable = [...]reflect.Kind{
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
	}
)

func isNillable(t reflect.Type) bool {
	for _, k := range nillable {
		if t.Kind() == k {
			return true
		}
	}

	return false
}

func NotNil[T any](v T) bool {
	typ := reflect.TypeOf(v)
	if !isNillable(typ) {
		return true
	}
	return !reflect.ValueOf(v).IsNil()
}

func Lt[T constraints.Ordered](limit T) func(T) bool {
	return func(v T) bool {
		return v < limit
	}
}

func Lte[T constraints.Ordered](limit T) func(T) bool {
	return func(v T) bool {
		return v <= limit
	}
}

func Gt[T constraints.Ordered](limit T) func(T) bool {
	return func(v T) bool {
		return v > limit
	}
}

func Gte[T constraints.Ordered](limit T) func(T) bool {
	return func(v T) bool {
		return v >= limit
	}
}

func Eq[T constraints.Ordered](limit T) func(T) bool {
	return func(v T) bool {
		return v == limit
	}
}

func ContainsText[T Text](substr string) func(T) bool {
	return func(v T) bool {
		return strings.Contains(string(v), substr)
	}
}
