package main

import (
	"fmt"

	"github.com/kami988/go-1.18-sample/1.18/11-proto/def"
)

func main() {
	a := def.GetReservedSpotUUIDByEventUUIDRequest{EventUUID: "id1"}
	fmt.Println(a.GetEventUUID())
}

// ファイル go.work を消してみよう
// go work init を実行してみよう
// go work -r . を実行してみよう

// spotSharedQueryService.proto を好きに変更してみよう
// make generate でpb.goファイルを更新しよう
// api/main.go で変更したpb.goを反映しているか確認してみよう
