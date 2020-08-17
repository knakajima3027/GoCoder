package main

import "fmt"

func main() {
	var card [4][13]int
	var n, number int
	var img, ans string

	fmt.Scan(&n)
	
	for i := 0; i < n; i ++ {
		fmt.Scan(&img, &number)

		if img == "S" {
			card[0][number - 1] = 1
		} else if img == "H" {
			card[1][number - 1] = 1
		} else if img == "C" {
			card[2][number - 1] = 1
		} else {
			card[3][number - 1] = 1
		}
	} 
		
	for i := 0; i < 4; i ++ {
		for j := 0; j < 13; j++ {
			if i == 0 {
				ans = "S %d\n"
			} else if i == 1 {
				ans = "H %d\n"
			} else if i == 2 {
				ans = "C %d\n"
			} else {
				ans = "D %d\n"
			}

			if card[i][j] == 0 {
				fmt.Printf(ans, j + 1)
			}
		}
	}

}