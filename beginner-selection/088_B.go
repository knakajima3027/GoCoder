package main

import (
	"fmt"
	
)

func main() {
	var N int
	var A []int

	fmt.Scan(&N)
	
	var tmp int
	for i:= 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}



	alice := 0
	bob := 0
	for i := 1; i <= N; i++ {
		if i%2 == 1 {
			bob += A[i]
		} else {
			alice += A[i]
		}
	}

	fmt.Println(bob - alice)
}