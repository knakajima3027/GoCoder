package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	fmt.Println(a*b, 2*a+2*b)
}