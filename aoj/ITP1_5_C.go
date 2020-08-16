package main

import "fmt"

func chessBord(h int, w int) {
	for i := 0; i < h; i++ {
		tmp := ""
		for j := 0; j < w; j++ {
			if i % 2 == 0 {
				if j % 2 == 0 {
					tmp += "#"
				} else {
					tmp += "."
				}
			} else {
				if j % 2 == 0 {
					tmp += "."
				} else {
					tmp += "#"
				}
			}
		}
		fmt.Println(tmp)
	}
}

func main() {
	var H, W int
	var ans[10000][2] int
	var n int

	for n = 0; n != -1; n++ {
		fmt.Scan(&H, &W)

		if H == 0 && W == 0 {
			break
		}
		ans[n][0] = H
		ans[n][1] = W
	}

	for i := 0; i < n; i++ {
		chessBord(ans[i][0], ans[i][1])
		fmt.Println("")
	}
}