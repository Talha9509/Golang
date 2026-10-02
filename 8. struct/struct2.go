package main

import (
	"fmt"
)

type vertex2 struct {
  x int
  y int
}

func main() {
	// using pointers and struct
  v := vertex2{1, 2}
	p := &v
  p.x = 23132
  fmt.Println(v)
}
