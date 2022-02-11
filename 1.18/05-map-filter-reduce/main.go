// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#map_reduce_filter
package main

import (
	"fmt"
)

func ForEach[T1 any](s []T1, f func(T1)) {
	for _, v := range s {
		f(v)
	}
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func main() {
	s := []int{1, 2, 3}

	// ForEach(s, func())

	floats := Map(s, func(i int) float64 { return 1.2 * float64(i) })
	fmt.Println(floats)

	sum := Reduce(s, 0, func(i, j int) int { return i + j })
	fmt.Println(sum)

	evens := Filter(s, func(i int) bool { return i%2 == 0 })
	fmt.Println(evens)
}

// 1.18 では標準packageとして搭載しない
