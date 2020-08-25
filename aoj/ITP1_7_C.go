package main

import "fmt"

func main() {
	var r, c, tmp, sum int
	var ans [101][101]int

	fmt.Scan(&r, &c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&ans[i][j])
			sum += ans[i][j]
		}
	}

	ans[r][c] = sum
	for i := 0; i < r; i++ {
		tmp = 0
		for j := 0; j < c; j++ {
			tmp += ans[i][j]
		}
		ans[i][c] = tmp
	}

	for i := 0; i < c; i++ {
		tmp = 0
		for j := 0; j < r; j++ {
			tmp += ans[j][i]
		}
		ans[r][i] = tmp
	}

	for i := 0; i < r + 1; i++ {
		for j := 0; j < c + 1; j++ {
			fmt.Print(ans[i][j])
			if j < c {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}