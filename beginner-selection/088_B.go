package main

import (
<<<<<<< HEAD
	"fmt"
	
)

func main() {
	var N int
	var A []int

	fmt.Scan(&N)
	
	var tmp int
	for i:= 0; i < N; i++ {
		fmt.Scan(&tmp)
		A = append(A, tmp)
	}



	alice := 0
	bob := 0
	for i := 1; i <= N; i++ {
		if i%2 == 1 {
			bob += A[i]
		} else {
			alice += A[i]
		}
	}

	fmt.Println(bob - alice)
}
=======
  "fmt"
  "sort"
  )

func main() {
  var N int
  var A []int

  fmt.Scan(&N)

  for i := 0; i < N; i++ {
    tmp := 0
    fmt.Scan(&tmp)
    A = append(A, tmp)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(A)))

  bob := 0
  alice := 0
  for i := 0; i < N; i++ {
    if i%2 == 0 {
      alice += A[i]
    } else {
      bob += A[i]
    }
  }
  fmt.Println(alice - bob)
}
>>>>>>> 3681d68c1563ae5abb7a5b9310ccd18ffb8f6743
