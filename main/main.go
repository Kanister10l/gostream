package main

import (
	"fmt"
	"github.com/Kanister10l/streams"
	"github.com/Kanister10l/streams/filters"
	"github.com/Kanister10l/streams/order"
)

func main() {
	ss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := streams.FlatMap(streams.Map(streams.OfSlice(ss).
		Filter(filters.Gt(2)).
		Order(order.Desc[int]),
		func(t int) int {
			return t * t
		}), func(t int) []int {
		return []int{t, t}
	}).Collect()

	fmt.Println(s)
}
