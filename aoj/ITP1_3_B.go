package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var answers[]string

	for i:= 0; i != -1; i++ {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		answers = append(answers, "Case " + strconv.Itoa(i+1) + ": " + strconv.Itoa(n))
	}

	for _, answer :=  range answers {
		fmt.Println(answer)
	}
}