package main

import "fmt"

func main() {
	var n, x int
	var ans [10000]int
	var num int

	for num = 0; num != -1; num++ {
		fmt.Scan(&n, &x)

		if  n == 0 && x == 0 {
			break
		}
		for i := 1; i < n + 1; i++ {
			for j := i + 1; j < n + 1; j++ {
				for k := j + 1; k < n + 1; k++ {
					if i + j + k == x {
						ans[num] += 1
					}
				}
			}
		}
	}

	for i := 0; i < num; i++ {
		fmt.Println(ans[i])
	}
}