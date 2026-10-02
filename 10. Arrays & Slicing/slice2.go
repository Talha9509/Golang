package main

import (
	"fmt"
)

func main() {
  a := []string{
    "a",
    "b",
    "c",
    "d",
    "e",
    "f",
  }
  w := a[:0]
  x := a[0:2]
  y := a[:3]
  z := a[4:]
  fmt.Println(a)
  printSlice(w)
  printSlice(x)
  printSlice(y)
  printSlice(z)
}

func printSlice(s []string){
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
