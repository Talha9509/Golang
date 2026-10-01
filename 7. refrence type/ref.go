package main

import (
	"fmt"
)

func main() {
  i := 2
  p := &i
  fmt.Println(i)
  fmt.Println(p)
  fmt.Println(*p)
  *p = 213 
  fmt.Println(i)
  fmt.Println(p)
  fmt.Println(*p)
}
