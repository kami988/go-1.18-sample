package main

import (
	"fmt"
)

func equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	fmt.Println(equal(12, 15))
	fmt.Println(equal("go", "go"))
}

// 7行目の comparable を any に置き換えてみよう
// 8行目の == を + に置き換えてみよう
// C++ではこの制約がない
