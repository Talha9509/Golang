package main

import "fmt"

var p = []int{1,2,3,4}

func main(){
  for i,v := range p {
    fmt.Printf("2**%d = %d\n", i,v)
  }
}
