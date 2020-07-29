package main

import (
  "fmt"
  "sort"
)

func main() {
  var N, tmp int
  fmt.Scan(&N)

  var d []int

  for i := 0; i < N; i++ {
    fmt.Scan(&tmp)
    d = append(d, tmp)
  }

  sort.Sort(sort.Reverse(sort.IntSlice(d)))

  res := 1
  for i := 0; i < N - 1; i++ {
    if d[i] > d[i+1] {
      res += 1
    }
  }

  fmt.Println(res)

}
