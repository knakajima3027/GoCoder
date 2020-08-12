package main

import "fmt"

func main() {
	var a, b, c, ans int;

	fmt.Scan(&a, &b, &c)

	ans = 0
	for i := a; i < b+1; i++ {
		if c% i == 0 {
			ans += 1
		}
	}

	fmt.Println(ans)
}