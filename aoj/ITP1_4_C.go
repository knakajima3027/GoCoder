package main

import "fmt"

func main() {
	var a, b, i int
	var ans [100000]int
	var op string
	
	for i = 0; i != -1; i++ {
		fmt.Scan(&a, &op, &b)

		if op == "?" {
			break
		} else if op == "+" {
			ans[i] = a + b
		} else if op == "-" {
			ans[i] = a - b
		} else if op == "*" {
			ans[i] = a * b
		} else {
			ans[i] = a / b
		}
	}

	for j:= 0; j < i; j++ {
		fmt.Println(ans[j])
	}
}