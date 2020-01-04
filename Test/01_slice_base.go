package main

import (
	// フォーマットの意味
	"fmt"
)

// 変数はvarを付けて宣言
// カッコでまとめた方がスマート
var (
	// これは初期化しているので型指定は要らない
	msg1 = "Hello"

	example = []string {
		"Golang",
		"hands-on",
		"in",
		"Osaka",
	}
)

// これは初期化していないので型指定が必要
var msg2 string

func main() {
	msg2 = "end"

	// 変数間にスペースが自動で付く
	fmt.Println(msg1, example, msg2);
}