package main

import "fmt"

func main() {
	var W, H, x, y, r int

	fmt.Scan(&W, &H, &x, &y, &r)

	flag := true
	if x-r < 0 || x+r > W {
		flag = false
	}
	if y-r < 0 || y+r > H {
		flag = false
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
