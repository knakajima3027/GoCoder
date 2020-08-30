package main

import "fmt"

func main() {
	var n, m, l int
	var A, B, ans [100][100]int

	fmt.Scan(&n, &m, &l)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < l; j++ {
			fmt.Scan(&B[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			tmp := 0
			for k := 0; k < m; k++ {
				tmp += A[i][k] * B[k][j]
			}
			ans[i][j] = tmp
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			fmt.Print(ans[i][j])
			if j != l - 1 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
}