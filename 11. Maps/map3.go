package main

import "fmt"

var mmm map[string]int

func main(){
  mmm = make(map[string]int)
  mmm["ans"] = 90
  fmt.Println(mmm["ans"])
  mmm["ans"] = 40
  fmt.Println(mmm["ans"])

  delete(mmm, "ans")
  fmt.Println(mmm["ans"])
  v, ok := mmm["ans"]
  fmt.Println(v, ok)
  mmm["ques"] = 40
  val, okk := mmm["ques"]
  fmt.Println(val, okk)

	for k, v := range mmm {
		fmt.Println(k, v)
	}
}