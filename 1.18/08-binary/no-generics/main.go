package main

func addInt(a, b int) int {
	return a + b
}

func addString(a, b string) string {
	return a + b
}

func main() {
	println(addInt(1, 2))
	println(addString("foo", "bar"))
}

// 下のコマンドを実行してみよう
// go build -o ./1.18/08-binary/no-generics-bin ./1.18/08-binary/no-generics/main.go
// ls -l ./1.18/08-binary/ | grep bin

// アセンブリを見たい方: objdump -d ./1.18/08-binary/no-generics-bin >> ./1.18/08-binary/no-generics/bin.txt
