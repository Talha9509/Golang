package main

import (
	"fmt"
)

type vertex3 struct {
  x int
  y int
}

func main() {
  v := vertex3{1, 2}
  x := vertex3{}
  c := vertex3{x: 1}
	a := &v
	b := &vertex3{6, 7}
  fmt.Println(v, x, c, a, b)
}
