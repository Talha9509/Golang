package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlicee(a)
	b := make([]int, 0, 5)
	printSlicee(b)

	c := b[:2]
	printSlicee(c)

}

func printSlicee(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
