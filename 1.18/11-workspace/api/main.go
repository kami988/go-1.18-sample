package main

import (
	"fmt"

	"github.com/kami988/go-1.18-sample/1.18/11-workspace/def"
)

func main() {
	a := def.Request{EventUUID: "id1"}
	fmt.Println(a.GetEventUUID())
}

// go work init を実行してみよう
// go work use -r . を実行してみよう

// spotSharedQueryService.proto を好きに変更してみよう
// make generate でpb.goファイルを更新しよう
// api/main.go で変更したpb.goを反映しているか確認してみよう
