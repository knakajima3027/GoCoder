package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, b, f, r, v int
	var ans [4][3][10]int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&b, &f, &r, &v)
		ans[b - 1][f - 1][r - 1] += v
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			tmp := ""
			for k := 0; k < 10; k++ {
				tmp += " " + strconv.Itoa(ans[i][j][k])
			}
			fmt.Println(tmp)
		}
		
		if i != 3 {
			fmt.Println("####################")
		}
	}
}