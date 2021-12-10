package streams

import (
	"fmt"
	"sort"
	"strings"
)

type Stream[T any] struct {
	values  []T
	onPanic func(interface{})
}

func (s Stream[T]) String() string {
	builder := strings.Builder{}
	builder.WriteString("Stream Of:")
	for _, value := range s.values {
		builder.WriteString(fmt.Sprintf("\n %v", value))
	}
	return builder.String()
}

func Of[T any](v ...T) Stream[T] {
	return OfSlice(v)
}

func OfSlice[T any](s []T) Stream[T] {
	var a Stream[T]
	a.values = s
	return a
}

func (s Stream[T]) Filter(f func(T) bool) Stream[T] {
	if s.onPanic != nil {
		defer onPanic(s.onPanic)
	}
	var a Stream[T]
	for _, v := range s.values {
		if f(v) {
			a.values = append(a.values, v)
		}
	}
	return a
}

func (s Stream[T]) Order(f func(T, T) bool) Stream[T] {
	if s.onPanic != nil {
		defer onPanic(s.onPanic)
	}
	sort.Slice(s.values, func(i, j int) bool {
		return f(s.values[i], s.values[j])
	})
	return s
}

func (s Stream[T]) Collect() []T {
	return s.values
}

func Map[T, U any](s Stream[T], f func(T) U) Stream[U] {
	if s.onPanic != nil {
		defer onPanic(s.onPanic)
	}
	a := Stream[U]{
		onPanic: s.onPanic,
	}
	for _, v := range s.values {
		a.values = append(a.values, f(v))
	}
	return a
}

func FlatMap[T, U any](s Stream[T], f func(T) []U) Stream[U] {
	if s.onPanic != nil {
		defer onPanic(s.onPanic)
	}
	a := Stream[U]{
		onPanic: s.onPanic,
	}
	for _, v := range s.values {
		a.values = append(a.values, f(v)...)
	}
	return a
}

func (s Stream[T]) OnPanic(f func(interface{})) Stream[T] {
	s.onPanic = f
	return s
}

func onPanic(f func(interface{})) {
	if r := recover(); r != nil {
		f(r)
	}
}
