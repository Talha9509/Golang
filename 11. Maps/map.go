package main

import "fmt"

type vertex struct {
	lat, long float64
}

var m map[string]vertex

func main(){
	m["a"] = vertex{34.3, 234.3}
	// m = make(map[string]vertex)
	// m["a"] = vertex{40.9, 32.3}
	// m["a"] = vertex{41.9, 72.3}
	fmt.Println(m["a"])
}