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

func main() {
	var data []string
	fmt.Println(data);

	for _, v := range example {
		data = append(data, v)
	} 

	// こちらの方がスマート
	//data = example

	fmt.Println(data);
}