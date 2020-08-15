package main

import "fmt"

func main() {
	var X, Y int
	ans := "No"

	fmt.Scan(&X, &Y)

	for i := 0; i <= X; i++ {
		if (2 * i + 4 * (X-i) == Y) {
			ans = "Yes"
		}
	}

	fmt.Println(ans)
}