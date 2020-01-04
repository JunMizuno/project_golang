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
	for _, v := range example {
		fmt.Println(v);
		if v == "in" {
			fmt.Println("◯");
		} else {
			fmt.Println("×");
		}
	} 
}