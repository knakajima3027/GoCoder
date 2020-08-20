package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
} 

func main() {
	var s string

	fmt.Scan(&s)

	ans := 0
	tmp := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' || s[i] == 'C' || s[i] == 'G' || s[i] == 'T' {
			tmp += 1
		} else {
			tmp = 0
		}
		ans = max(ans, tmp)
	}
	
	fmt.Println(ans)
}