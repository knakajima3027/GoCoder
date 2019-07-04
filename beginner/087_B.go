package main

import "fmt"

func main() {
	var A, B, C, X int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&X)

	res := 0 
	for i := 0; i <= A; i ++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				if 500 * i + 100 * j + 50 * k == X {
					res += 1
				}
			}
		}
	}

	fmt.Println(res)
}