package main

import (
	"errors"
	"testing"
)

func FuzzCalc(f *testing.F) {
	f.Add(1, 2, "+")
	f.Fuzz(func(t *testing.T, v1, v2 int, ope string) {
		_, _ = Calc(v1, v2, ope)
	})
}

func Calc(v1, v2 int, ope string) (int, error) {
	switch ope {
	case "+":
		return v1 + v2, nil
	case "-":
		return v1 - v2, nil
	case "*":
		return v1 * v2, nil
	case "/":
		return v1 / v2, nil
	}
	return 0, errors.New("")
}

// 下のコマンドを実行してみよう
// go test ./1.18/10-fuzz/ -fuzz=FuzzCalc -fuzztime=30s

// 1.18でf.Fuzzに登録出来る型
// []byte
// string
// bool
// byte
// rune
// float32
// float64
// int
// int8
// int16
// int32
// int64
// uint
// uint8
// uint16
// uint32
// uint64
