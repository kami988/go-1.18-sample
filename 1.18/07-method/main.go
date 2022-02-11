package main

import "fmt"

type List[T any] struct {
	Next  *List[T]
	Value T
}

func (t *List[T]) add(newList List[T]) {
	list := t
	for {
		if list.Next != nil {
			list = list.Next
			continue
		}
		list.Next = &newList
		break
	}
}

func (t *List[T]) printFull() {
	list := t
	for {
		if list.Next != nil {
			fmt.Print(list.Value, " -> ")
			list = list.Next
			continue
		}
		fmt.Print(list.Value, " -> ")
		break
	}
	fmt.Println("nil")
}

func main() {
	intList := List[int]{Value: 10}
	intList.add(List[int]{Value: 20})
	intList.add(List[int]{Value: 30})
	intList.printFull()

	stringList := List[string]{Value: "a"}
	stringList.add(List[string]{Value: "b"})
	stringList.add(List[string]{Value: "c"})
	stringList.printFull()

	fn := func() { fmt.Println() }
	fn()
}

// 10行目の add に型パラメータを導入しようとすると？
// 48行目の funcに型パラメーターを導入しようとすると？
