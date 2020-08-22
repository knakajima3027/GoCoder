package main

import "fmt"

func main() {
	var f, m, r, i int
	var ans [50]string

	for i = 0; i != -1; i++ {
		fmt.Scan(&f, &m, &r)
		if f == -1 && m == -1 && r == -1 {
			break
		}

		sum := 0
		if f != -1 {
			sum += f
		}

		if m != -1 {
			sum += m
		}

		if f == -1 || m == -1 {
			ans[i] = "F"
		} else if sum >= 80 {
			ans[i] = "A"
		} else if sum >= 65 {
			ans[i] = "B"
		} else if sum >= 50 {
			ans[i] = "C"
		} else if sum >= 30 {
			if r >= 50 {
				ans[i] = "C"
			} else {
				ans[i] = "D"
			}
		} else {
			ans[i] = "F"
		}
	}

	for j := 0; j < i; j++ {
		fmt.Println(ans[j])
	}
}