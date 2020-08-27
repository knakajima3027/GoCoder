package main

import "fmt"


func main() {
	var N, tmp uint64
	var A [5]uint64

	fmt.Scan(&N)
	fmt.Scan(&A[0])
	fmt.Scan(&A[1])
	fmt.Scan(&A[2])
	fmt.Scan(&A[3])
	fmt.Scan(&A[4])

	tmp = 10000000000000000
	for i := 0; i < 5; i++ {
		if tmp > A[i] {
			tmp = A[i]
		}
	}

	if N % tmp == 0 {
		fmt.Println(N / tmp + 4)
	} else {
		fmt.Println(N / tmp + 5)
	}
}