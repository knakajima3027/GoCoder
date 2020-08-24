package main

import (
	"fmt"
	"math"
)

func inThree(x int) bool {
	if x % 3 == 0 {
		return true
	}

	for i := 0; i < 10; i++ {
		tmp := x
		tmp = tmp / int(math.Pow(10.0, float64(i)))

		if tmp == 0 {
			break
		}
		
		if tmp % 10 == 3 {
			return true
		}
	}

	return false
}

func main() {
	var n int

	fmt.Scan(&n)

	for i := 1; i < n + 1; i++ {
		if inThree(i) {
			fmt.Print(" ")
			fmt.Print(i)
		}
	}
	fmt.Println()
}