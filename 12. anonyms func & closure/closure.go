package main

import (
  "fmt"
)

func adder() func(int) int {
  sum := 0
	// the inner func will have the access to outer func variables even after inned func returns. so it can use it as state
  return func(x int) int {
    sum += x
    return sum
  }
}

func main(){
  pos, neg := adder(), adder()
	// here in the func pos and neg, it will remember the sum and update it. it is due to closures
  for i := 0; i < 10; i++ {
    fmt.Println(pos(i), neg(-2*i))
  }
}