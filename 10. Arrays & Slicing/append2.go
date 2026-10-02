package main

import "fmt"

func main(){
  var a []int
  for i := 0; i < 40; i++ {
    a = append(a, i)
    // fmt.Println(a)
    fmt.Printf("length: %d, capacity: %d\n", len(a), cap(a))
  }
}
