package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func main() {
	a := 1
	b := 2
	fmt.Println(sum(a, b))
}
