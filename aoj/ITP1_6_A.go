package main

import "fmt"

func main() {
	var n int
	var A [100]int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	
	for i := 0; i < n; i++ {
		if i != n - 1 {
			fmt.Printf("%d ", A[n - i - 1])
		} else {
			fmt.Printf("%d\n", A[n - i - 1])
		}
	}
}