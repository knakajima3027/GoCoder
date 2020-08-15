package main

import (
	"fmt"
	"math"
)

func main() {
	var X, n, p int
	var P [103]int
	var tmp float64
	tmp = 10000
	ans := 0

	fmt.Scan(&X, &n)

	for i := 0; i < n; i++ {
		fmt.Scan(&p)
		P[p+1] = 1
	}

	for i := 0; i < 103; i++ {
		if P[i] == 0 {
			if float64(tmp) > math.Abs(float64(X - i + 1)) {
				ans = i - 1
				tmp = math.Abs(float64(X - i + 1))
			}
		}
	}
	fmt.Println(ans)
}