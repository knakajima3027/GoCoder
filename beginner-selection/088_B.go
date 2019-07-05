package main

import (
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
