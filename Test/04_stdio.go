package main

import (
	"fmt"
)

var (
	example = []string {
		"Golang",
		"hands-on",
		"in",
		"Osaka",
	}
)

var data []string

func main() {
	// 文末のモノで処理終わりを判定しているので、ヘタなインデントなどはエラーの元になってしまう
	data = example
	for _, v := range data {
		fmt.Println(v);
		var ans string
		fmt.Print("入力>？")
		fmt.Scanln(&ans)
		if ans == v {
			fmt.Println("◯")
		} else {
			fmt.Println("×")
		}
	} 
}