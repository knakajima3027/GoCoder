
package main

import "fmt"

func main() {
  var s string
  dream := [4]string{"maerd", "remaerd", "esare", "resare"}

  fmt.Scan(&s)

  tmp := ""
  for i := 0; i < len(s); i++ {
    tmp += s[len(s)-i-1:len(s)-i]

    for j := 0; j < 4; j++ {
      if dream[j] == tmp {
        tmp = ""
      }
    }
  }

  if tmp == "" {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }

}

