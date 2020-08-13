package main

import "fmt"

func main() {
  var a, b int
  var ans1, ans2 int
  var ans3 float64

  fmt.Scan(&a, &b)
  
  ans1 = a / b
  ans2 = a % b
  ans3 = float64(a) / float64(b)

  fmt.Printf("%d %d %f\n", ans1, ans2, ans3)
}
