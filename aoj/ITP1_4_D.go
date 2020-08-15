package main

import "fmt"

func searchMin(A [10000]int, n int) int {
	ans := 10000000
	for i := 0; i < n; i++ {
		if ans > A[i] {
			ans = A[i]
		}
	}
	return ans
}

func searchMax(A [10000]int, n int) int {
	ans := -1000000
	for i := 0; i < n; i++ {
		if ans < A[i] {
			ans = A[i]
		}
	}
	return ans
}

func seekSum(A [10000]int, n int) int {
	ans := 0
	for i := 0; i < n; i++ {
		ans += A[i]
	}
	return ans
}

func main() {
	var n, min, max, sum int
	var A [10000]int

	fmt.Scan(&n)
	
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	
	min = searchMin(A, n)
	max = searchMax(A, n)
	sum = seekSum(A, n)

	fmt.Println(min, max, sum)

}