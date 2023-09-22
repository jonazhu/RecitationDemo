package main

import (
	"fmt"
)

func main() {
	PrintAllNum(100)
}

func PrintAllNum(n int) {
	for i := range n {
		fmt.Println(i)
	}
}