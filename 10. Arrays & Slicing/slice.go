package main

import (
	"fmt"
)

func main() {
  var a = [4]string{
    "a",
    "b",
    "c",
    "d",
  }
  x := a[0:2]
  y := a[1:3]
  fmt.Println(a,x,y)
	y[0] = "4"
  fmt.Println(a,x,y)

	qw := []struct{
    i int
    b bool
  }{
    {2, true},
    {3, true},
    {5, false},
  }
  fmt.Println(qw)

}