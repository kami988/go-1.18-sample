package main

import (
	"fmt"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func EqualInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// slices
	is1 := []int{1, 2, 3}
	is2 := []int{1, 2, 4}
	fmt.Println("Equal:   ", slices.Equal(is1, is2)) // false

	ss1 := []string{"a", "b", "c"}
	ss2 := []string{"a", "b", "c"}
	fmt.Println("Equal:   ", slices.Equal(ss1, ss2)) // true

	fmt.Println("Contains:", slices.Contains(ss1, "b"))
	fmt.Println("Index:   ", slices.Index(ss1, "b"))

	fmt.Println("Insert:  ", slices.Insert(ss1, 2, "g"))
	fmt.Println("Delete:  ", slices.Delete(ss1, 2, 3))

	// maps
	m := map[string]bool{
		"SpotOwn":  false,
		"SpotEdit": true,
	}
	fmt.Println("Keys", maps.Keys(m)) // 順不同
}

// slices, maps package は1.19で標準packageとして搭載予定
// https://github.com/golang/go/issues/45955
// https://github.com/golang/go/issues/47649
