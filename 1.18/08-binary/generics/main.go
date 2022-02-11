package main

type Addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128 | string
}

func addddd[T Addable](a, b T) T {
	return a + b
}

func main() {
	println(addddd(1, 2))
	println(addddd("foo", "bar"))
}

// 下のコマンドを実行してみよう
// go build -o ./1.18/08-binary/generics-bin /repo/1.18/08-binary/generics/main.go
// ls -l ./1.18/08-binary/ | grep bin

// アセンブリを見たい方: objdump -d ./1.18/08-binary/generics-bin >> ./1.18/08-binary/generics/bin.txt
