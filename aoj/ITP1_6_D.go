package main

import "fmt"

func main() {
	var n, m int
	var arr1 [100][100]int
	var arr2 [100]int

	fmt.Scan(&n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&arr1[i][j])
		}
	}
	
	for i := 0; i < m; i++ {
		fmt.Scan(&arr2[i])
	}

	for i := 0; i < n; i++ {
		tmp := 0
		for j := 0; j < m; j++ {
			tmp += arr1[i][j] * arr2[j]
		}
		fmt.Println(tmp)
	}
}
