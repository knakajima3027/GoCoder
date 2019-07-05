package main

import (
  "fmt"
)

func main() {
  var N, Y int
  fmt.Scan(&N, &Y)

  flag := true
  for i := 0; i <= N; i++ {
    for j := 0; j <= N; j++ {
      th := (Y - 10000*i - 5000*j)/1000
      if (Y - 10000*i - 5000*j)%1000 == 0 && i + j + th == N  && th >= 0 {
        fmt.Println(i, j, th)
        flag = false
        break
      }
    }
    if !(flag) {
      break
    }
  }

  if flag {
    fmt.Println(-1, -1, -1)
  }

}
