package main

import (
	"fmt"
	"math"
)

func main() {
	var r, S, H float64
	pi := math.Pi

	fmt.Scan(&r)

	S = pi * r * r
	H = 2 * pi * r	
	
	fmt.Printf("%f %f\n", S, H)
}