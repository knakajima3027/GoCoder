package main

import "fmt"


func main() {
	var n int

	fmt.Scan(&n)

	var A[2000] int

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	flag := false
	res := 0
	for i := 0; true; i++ {
		for i := 0; i < n; i++ {
			if A[i]%2 == 0 {
				A[i] /= 2
			} else {
				flag = true
			}
		}

		if flag {
			break
		}
		res += 1

	}

	fmt.Println(res)
}