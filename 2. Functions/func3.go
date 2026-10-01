package main

import "fmt"

func split(sum int)(x, y int){
	// here x and y are initialized with 0 like 
	// x = 0
	// y = 0
	x = sum * 4 / 9
	y = sum - x
	return
	// here it will return in format x, y int
}

func main(){
	fmt.Println(split(17))
}
