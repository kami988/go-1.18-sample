package main

import (
	"fmt"
)

type addable interface {
	int | int32 | int64 | float32 | float64 | complex64 | complex128 | string
}

func add[T addable](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1., 0.18))

	fmt.Println(add("foo", "bar"))
}

// 11行目の add の addable を any に換えてみよう
