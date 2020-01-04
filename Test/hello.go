package main

import (
	"fmt"
)
var (
	msg1 = "Hello"
	msg2 string
	msg3 string
	msg4 string
)
func main() {
	msg2 = "world"
	msg3 = "!"
	msg4 = "!!!"
	// 変数間にスペースが自動で付く
	fmt.Println(msg1, msg2, msg3, msg4);
}