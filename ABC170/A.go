package main 

import "fmt"

func main() {
	var x [5]int

	for i := 0; i < 5; i++ {
		fmt.Scan(&x[i])
	}

	for i := 0; i < 5; i++ {
		if x[i] == 0 {
			fmt.Println(i + 1)
		}
	}
}