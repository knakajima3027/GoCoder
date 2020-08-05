package main

import "fmt"

func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func max(x int, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	var a, b int
	var answers [30000][2]int
	var i int
	for i = 0; i != -1; i++ {

		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		answers[i][0] = min(a, b)
		answers[i][1] = max(a, b)
	}

	for j := 0; j < i; j++ {
		fmt.Println(answers[j][0], answers[j][1])
	}
}
