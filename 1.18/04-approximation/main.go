package main

import (
	"fmt"
)

type (
	intOrFloat interface {
		int | float32 | float64
	}
	// Minute は vo です
	Minute int
)

func doSomething[T, U intOrFloat](x, y T) T {
	switch {
	case x > y:
		fmt.Println(x)
		return x
	default:
		fmt.Println(y)
		return y
	}
}

func main() {
	// cancelMin1 := Minute(15)
	// cancelMin2 := Minute(30)
	// doSomething()
}

// 18, 19, 20行目のコメントアウトを解除して、doSomethingの引数にcancelMinを入れてみよう
// 7行目の各型名の前に~を追加してみよう
// 13行目のintOrFloatをconstraints.Orderedに置き換えてみよう
