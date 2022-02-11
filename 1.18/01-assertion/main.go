package main

import "fmt"

func print[T any](x T) {
	fmt.Println(x)
}

type mySt[X any, Y any] struct {
	x X
	y Y
}

func main() {
	print[int64](10)
	print("hello")

	a := mySt[int, string]{x: 10, y: "hello"}
	fmt.Println(a)
}

// 15行目 int64 を int に置き換えてみよう
// 18行目 [int, string]を消してみよう
